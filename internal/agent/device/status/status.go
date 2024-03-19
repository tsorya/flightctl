package status

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/flightctl/flightctl/api/v1alpha1"
	"github.com/flightctl/flightctl/internal/client"
	"github.com/flightctl/flightctl/internal/tpm"
	"github.com/flightctl/flightctl/pkg/executer"
	"github.com/sirupsen/logrus"
	"k8s.io/klog/v2"
)

var _ Manager = (*StatusManager)(nil)

// NewManager creates a new device status manager.
func NewManager(
	deviceName string,
	tpm *tpm.TPM,
	executor executer.Executer,
) *StatusManager {
	exporters := []Exporter{
		newSystemD(executor),
		newContainer(executor),
		newSystemInfo(tpm),
	}
	return &StatusManager{
		deviceName: deviceName,
		exporters:  exporters,
		conditions: DefaultConditions(),
	}
}

// Collector aggregates device status from various exporters.
type StatusManager struct {
	deviceName       string
	managementClient *client.Management
	exporters        []Exporter
	log              *logrus.Logger
	logPrefix        string
	conditions       *[]v1alpha1.Condition
}

type Exporter interface {
	Export(ctx context.Context, device *v1alpha1.DeviceStatus) error
}

type Collector interface {
	Get(context.Context) (*v1alpha1.DeviceStatus, error)
}

type Manager interface {
	Collector
	Update(context.Context, *v1alpha1.DeviceStatus) error
	SetClient(*client.Management)
	UpdateConditionError(ctx context.Context, reason string, err error) error
	UpdateCondition(ctx context.Context, conditionType v1alpha1.ConditionType, conditionStatus v1alpha1.ConditionStatus, reason, message string) error
}

func (m *StatusManager) SetClient(managementCLient *client.Management) {
	m.managementClient = managementCLient
}

func (m *StatusManager) Get(ctx context.Context) (*v1alpha1.DeviceStatus, error) {
	deviceStatus, err := m.aggregateDeviceStatus(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get device status: %w", err)
	}

	return deviceStatus, nil
}

func (m *StatusManager) aggregateDeviceStatus(ctx context.Context) (*v1alpha1.DeviceStatus, error) {
	deviceStatus := v1alpha1.DeviceStatus{}
	for _, exporter := range m.exporters {
		err := exporter.Export(ctx, &deviceStatus)
		if err != nil {
			klog.Errorf("failed getting status from exporter: %v", err)
			continue
		}
	}

	// add conditions
	deviceStatus.Conditions = m.conditions

	return &deviceStatus, nil
}

func (m *StatusManager) Update(ctx context.Context, status *v1alpha1.DeviceStatus) error {
	if m.managementClient == nil {
		return fmt.Errorf("management client not set")
	}

	// we keep our status conditions in memory, so don't stomp on it
	if status.Conditions == nil {
		return fmt.Errorf("status conditions not set")
	}

	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(status)
	if err != nil {
		return fmt.Errorf("failed to encode device: %w", err)
	}
	if err := m.managementClient.UpdateDeviceStatus(ctx, m.deviceName, buf); err != nil {
		return fmt.Errorf("failed to update device status: %w", err)
	}
	// update conditions
	m.conditions = status.Conditions
	return nil
}

func (m *StatusManager) UpdateCondition(
	ctx context.Context,
	conditionType v1alpha1.ConditionType,
	conditionStatus v1alpha1.ConditionStatus,
	reason,
	message string,
) error {
	if m.managementClient == nil {
		return fmt.Errorf("management client not set")
	}

	status, err := m.Get(ctx)
	if err != nil {
		return err
	}

	if status.Conditions == nil {
		return fmt.Errorf("status conditions not set")
	}

	if SetProgressingCondition(status.Conditions, conditionType, conditionStatus, reason, message) {
		// log condition change
		m.log.Infof("%sSet progressing condition: %s", m.logPrefix, reason)
	}

	return m.Update(ctx, status)
}

func (m *StatusManager) UpdateConditionError(ctx context.Context, reason string, serr error) error {
	if m.managementClient == nil {
		return fmt.Errorf("management client not set")
	}

	status, err := m.Get(ctx)
	if err != nil {
		return err
	}

	if status.Conditions == nil {
		return fmt.Errorf("status conditions not set")
	}

	if SetDegradedConditionByError(status.Conditions, reason, serr) {
		// log condition change
		m.log.Infof("Set degraded condition by error: %v", err)
	}

	return m.Update(ctx, status)
}
