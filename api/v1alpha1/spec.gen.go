// Package v1alpha1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1alpha1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+x9/XPcNpLov4Ka26rs5kYj25vb2lXV1StFdhK/+EMlybl6b+V3BZE9MzhxAAYAJc+m",
	"9L+/QgMgQRKcIUeftvhLIg/x0Wg0Gv2NPyaJWOWCA9dqcvDHRCVLWFH88zDPM5ZQzQQ/1VQX+GMuRQ5S",
	"M8B/cboC8/8UVCJZbppODia/FCvKiQSa0osMiGlExJzoJRBajTmbTCd6ncPkYKK0ZHwxuZlOTKd1e8Sz",
	"JRBerC5AmoESwTVlHKQi10uWLAmVgNOtCeM9p1GaSrvi+kwfyll8GyIuFMgrSMlcyA2jM65hAdIMr0p0",
	"/UnCfHIw+bf9Csv7DsX7LfyemYFuELzfCyYhnRz806LYIyaAvJzlcwmBuPgfSLQBID70wR8T4MXKjHos",
	"IaeIjenk1Axo/zwpOLd/vZFSyMl08olfcnHNJ9PJkVjlGWhIgxkdRqeTL3tm5L0rKg28ykzRgiGcs/Ux",
	"AKL1rYKq9cmD2fpQwd36FCykjip1WqxWVK67qJ3xudhK7aaRXOF4JAVNWcb4Askmo0oTtVYaViEJES0p",
	"V6yTVgcTU30ZUaLqRzqRgQIS+gVoppeGJl/DQtIU0gjZDCaV+pzVHJ1Ngsk720SopN6gBNcgoNDLI8Hn",
	"bNHea/PNsJ85W5i9qpMHLfTSIynSDfEQ2V/T7dPJu45e5kurU2M3y4mrwWI7e3T86QSUKGQC7wVnWsjT",
	"HBKEPMs+zicH/9xMYrHONwZjRwYHc4NYOGULc1RP4PcClG6vqbMpkZBLUGZCQol0PxqOS4liCw4pSaq+",
	"ZC7FCg/V0WF7H3L2G0iFE7ZwevzWfSMpzBkHhaNc2d8gJXax9rpiqoLKHlUxJ5QTi9IZOTXXglRELUWR",
	"pYYurkCalSRiwdm/ytEU0cJxAG1WZW4KyWlGrmhWwJRQnpIVXRMJZlxS8GAEbKJm5L2QlrcckKXWuTrY",
	"318wPbv8u5oxYXZrVXCm1/vmbpTsotBCqv0UriDbV2yxR2WyZBoSXUjYpznbQ2A5noTZKv036fZWxSj0",
	"kvG0jcpfGU8JM7tlW1pQK4x5tnfy5vSM+PEtVi0Cgy2vcGnwwPgcpG1Z7jPwNBeMa/xHkjHgmqjiYsW0",
	"8tRi0DwjR5RzockFkCJPqYZ0Rt5yckRXkB1RBfeOSYM9tWdQFsXlCjRNqabb+PlHRNF70BTvAHdQN/Xo",
	"PFr2oPa9SLqHsd1bzKc6bY5SgkU6yKPcqGued2wQ4zDNLRlm5i8xJ93saOQU98wpmIZVRKh+t21nzGVa",
	"9t2JOs3sDhwqJV2PfOtx+JbZasu1hvEJu/uDGIWXXurb+1+S5jlIQqUoeEooKRTIvUSCwSk5Oj2ZkpVI",
	"IYOUCE4uiwuQHDQowgTikuZsFkgaanb1crYZhCZXgS85k1bfgEQYfLaAdN0hJWkhS4ZxRTOWMr0uFc0A",
	"jsl0YvUKq2n+9VVU8YQvWuIW0TRFjYJmx3UVxh+y1gY3D08d4DdmYEK1pSxQXp83yCV6STXxGEahzGA5",
	"F3mR4U8Xa/z18PgtQU1aGsxje7Nww9PYalVooz5NIgQgu4TJsyWQC6rgbz/sAU9ECik5fvO++vvXo9N/",
	"e/nCQDMj76lOlo6HmztpVoqYDLKUME5oSAyb5FTLEcINuVjrqGiPgqv8EDWSvOWpJTAESZYEYftYVo9c",
	"6veCZmzOICXOFNCapmARNvfp7ev736QABkUXEKH0T/g7otwsAtku4GVwCWtiewWrd/YbplRRl/hrN8RW",
	"4jUrjtumPgTGqPvHS4MHylIOCShjGM8rZbguaqJ5LsUVzfZT4Ixm+3PKskICsdKfXzou0gDvbGkqgnaj",
	"ZzEjxqwJfGEKbU51Thfyp+jpdAO2FbhphTUieAIVwvucK8NVkb1FMHFUfrNGFrOrIjxjM/Kr0fVJEjSU",
	"QA4Rb5BOyWvgzPzfoOcnyjKEqaS9frpyCcXk5rPhpXNaZIaD3bSItUEiwdKihFGO273wak+t/UnhfSI4",
	"EGqOofY0kBRSojiizU57OdYQutf02zaOjCp9VtqrztiqY+PR1qXZCuxMJWiVrQtSKyQZuBxtakEoF3oJ",
	"chZSgZGG9sxYcblEGR6y1Szn2hFmD4oR8jx26IUotIN4synOW4J/Bg722o6vfuYFm9mibGkZTR0b11Qh",
	"NzSXWEqK3E4b3vN/+yF6z0ugKjb5ny8kg/lfiP1eyRF+xu9Ur3X21BT9qF4z9CP17Ba1TDormYNgGiO4",
	"cvnV7m88KhXP9KbLM1mYYX6imYLBxsrGuG6sxq9+6MbPoZ2xjocAOs+JrMHS/2m5EkLtWNJhkoBSzF48",
	"tX/483tMpcKmp2ue4B8fr0BmNM8ZX5xCBolREibTyW9G8jSYMKqH8wrkkPif3xeZZnkGH685BO374esN",
	"lyLLVsC1u8OCRXXec33alBjpbFGi6gRyoZgWch3Fk0FP54cWMsOPJWJ/ygB0B3bxm8fla7hiCQSItj+E",
	"6La/tJB+BqvcXJFOjXJ7YCipUFqs7t62O22yl1MrxTm/heEuK9vesNMEoSjlYzVry/IGWLu4Nuuyv9fN",
	"wPlyrVhCM5Lix9lowBlNvaOpV+1XLKP/be367GDEjV2udrSaP63Daeox0GWRGOQhbxsm3tPcHNWIW9Wi",
	"JcqHphNlvX87e1VbGPTmbjduN86sa7ELWxJ4ChLSTq7mWZqT4VPPNW23wDe5TROtz7MRXiUyaIO6ODk+",
	"euOOalQpV+Y+Ffzt68jXBji1scKe3XD9IsSl8pdc41aYa5AncCEEXrFt1cB0JfAFksJo+NicSN+eAEeN",
	"wd1nNHE6omGBRgJ34vw100uCyoojPnXOhUQbATO3HzlbgoKyu0iSQrqpgo1bUuVmRo0zy8S1AcFcrblQ",
	"es9+I5qqSzU7533N5BZFFgVmtZ5VNO0kCE8pi/RDVOGa3z+eLDF7A2mypHwBiizpFZALAN7U752QMBRL",
	"uHzYhKULmAsJ/QnKtg8oCvcVN/U+kOWmC6iKVUR1D0Rj5+tNNQ68kmweBBlx0qESHohobjr51ltcIdOd",
	"UUY9r6boaO6Oasf7bL2WOga6fQyUta6U8U/Mz3M3NohNwA+NfNo6Vhg/R5Wqa+NVwNknroo8F7J/qFx0",
	"5nKK6Ndy3ujXCpiOzwGE5crjjvfqW93Lbn9Xo0722E71YCMGMLDRX/6t+svt/n48jcvGbBW1lgulJQDB",
	"ry5gW5JPJ++2axJ2wI2AdIXTxkFpaDgfTy1Ut4ekIdm01Yakw6V0tqykDE0vgXspw3AOK6o6/dNKXVbQ",
	"8M6CGXlDk6UbwBykUjJyzk4hU6sUrLGfZYxp7/NsFnSYWF/TFjd+REnzXtItEaZJt0vKI9cZRTs2O8mL",
	"vvJnOJC9w6eTlKnL2/RfwUr0laliIzS9dHkxKQd10PXFTXcc9X9R6eLcjyTTLKHZzhHVsYnDgO3212ry",
	"2NcAoNhnD2TsW+jzCAxXbQrpiLn295z9XrcvV1cjM11WjFMtZDD22gZfuME9NQgOPWziPzNtjTXHUlyx",
	"FCqr+KZev5ZBJKeQSNCDOr/lGeOww6y/aJ3HusWIsskiqnSa9qasqE6Wx1QbiakeQJTbHycHk//3T7r3",
	"r8/mPy/2/rH337PP3/8pdiNuV5CWRnHsd0Yr64/Zzp6d3J1o83+cvNUWUQ18Lv/HijLOx1HXKfvLWw3X",
	"SmwH7O2RDkH/in55B3yhl5ODV//xt2lzOw73/u+LvX8cnJ/v/ffs/Pz8/PsdN6Vbj+2KRAm/ht6cuE5Y",
	"RaVQr4oT19dIiFpSltmcq0QXNKtiFegGn1Bls+1HFxEzdv8gk3KJ9i7HS586i4QBMxppEULfL+61iieJ",
	"HmDHObevtWZ+NsqF1wx30rTNCEatPwVA8aJfzMaA81rOUjuxQ+/wARZ/R751W78/oW+d8aPHAFX7m+nE",
	"aShDTEtph58hoMoaVNM63YcICze5JBbchQqyCj/BhnZLNA+QKOdMmT406O6MRbfKjusaIpDnPuIdHk+L",
	"q2zI08mxuAYJ6cf5fEfprgZFMGvrWwBI5Gtddqt9CsGNfK6tIPI9IvnVjlH04ihbOJuDDRZlqdovCpai",
	"Lafg7PcCsjVhqVHI5+vQktu+DwJFPq7bHQYtDD9Hw5gP/KyGbVGdQY71btXH/FEITd6+HjKUARjN43b9",
	"cTg/+kbk1KubPSdoqnMhSsp1tKHoPgEN+/eOurRAdZpcL4GXgdk21HnOMiAOHB+h+VUr1Ebp+IlZP2ov",
	"KEzjjx4BMUByaoS/GH7NF4NcL7iir8W5QBhv+EYMptGXwpTtmFBOnAlOEGDof6F+axK3M5JQTszhM/hl",
	"EkOd1j0Ib6sdoX773bn7wd0q9tq7y1ulBvdut0p7iOBW+ZSfidc2D+RjoT/O3d9BHNkuV0htymCKyNdw",
	"1mjnRkBb/WvrJgg9TA0FjDhRpB7joPzpnmcAmkjQheSQWuYxB50s0blIFOOLDAjG3LUvA9UUXLrCUtox",
	"t00oLyTQy1Rc841wXqzJuZ/1fOLEmWhIihaaZvEDjZ+CAhixmeKVKCyhP/BynVC5abnNqF9c+7SxPQ3w",
	"oyeHqcvHDnxMmbq0mSxteutm0iXXjLLr+pibmSrO8TkabNmKvW3D0mqyIRPfpZbgxYDdNqrhoxtwDM18",
	"dqGZreM0LEqz3f1us+47gvHtxdOyo9kQ/BbN+S8+mQaUuYVRUAzyrDC2zcdpYfuAlV0IkQFFudZ/PdTd",
	"Mx1i8IEZHHOKqHaFn8LprqmqzdTPFuV7/Ljunv3HtZ+9UcrKfJVR0TGjF5D1uXGrLvW57QA1Hdj9pAXG",
	"MK0bEUxb79hyP3vRRTwYJNqsHhfSajJeDY8dIRLdkl5KaFt+GMNGvtGwkfjFtZ0DmGZ2n4OG1sfRavud",
	"IprKBThPSJszJEq2p0yUtBPEkvvDolDKJn+Vib4xBKcN51X/lIk7YOqHTVbuU0JdTCu5Zkamrrg7U95i",
	"gWquoWYokYpIqfLkNnN/g9l+297h1+toOMzF1+tyqASSQayplGRuppsT00OSadFVO1V9NjgDvZ1XDbfg",
	"wRtcf8Nyx62hv+1j7sqgxvY+cXqrFlqm4t5MJ3XbZdyksc4RN6WN1x4GI8SV9TaFsyywDDfBm8KOsFgD",
	"elNW4qq0gkFPy1cNuHKs2q/lwLVf/Sw3LqOzvbCfnGkq0J/dgU/HaNlRTR7V5Mq7YU7KMNXYdrlbdRjH",
	"jKs65ae6eoM/j+f40XWaah/6edOQYY/KyzeqvFTsJH6ONygp6LfYqpgoV85h69KMYO9rPyC9uaINMbnr",
	"IdLDm07IOCdsumM80N247tASgo/DNAPrw+ob+4etpwTQK02zbE1Y5RWrWthcUXNkMMY18SW+VpTTBaAO",
	"5TUvLAB3vXSiZisyeZiwXzrkbh/bl7a8pdt3vjMNcIuCgGWNWOJCBP1pGhRZHQvp9q7sHXMPgkFclw2w",
	"n0AuSgdgVEmf00xBE9A+dYv80H6phezw1v45F1hIxtytK6HhLxinY8vP9KrnbUZ2baJLjcal9/Z4tnf5",
	"ZtrK7Wf6xIzwR4c7M/Kggl9hx+MNgeE8wEZ1dQpSKCDUFWhc84TYL5iY2w5bRmZ9AldMxeNwWuUOSvBa",
	"naddDtRmjQKLk7ijNYgZOvgjyC9oVgaFxJUK7B2D9KbsE2XowZCf2/sYBJb3m80GfqXxu8MN9jmaVRCD",
	"uE1AwK9+ozIWds+JyO1pLWXtX9/8n//87fDdpzckp0yiQGvUaaoI8CsmBUcOfkUlM5OpstxZhZNhVSNl",
	"0WGsMIKTkZe1MLKXDzebEsaTrEgx7oSvCZWLYoXXXaHMb0pTnlKZErWELDNErekXF2llq466LFxFVq7W",
	"k59JkZzlmKS+QL/a1CyazW1M2zXICghS8BQDtC6oWpK9BG86+BI3fl4LefmayW0BCYwH7rUKmdY6eQFE",
	"FtwKr2xOGOpHGcw1gVWu1+YHbFc28pU2FVmK1aBoMbMffUltGA8MCL5Xdk2MthvnPh4HqdkKRNFRu3ZF",
	"v7BVsapqAGNphPChGRviqIWhC3yzZEbOOW6W7+LUwIsweJJi7SzD8NgVEBfTQ875XLjxL9aEWk+qUQdm",
	"5NRng1c/YsjlwTnfI9+p7xAgZYsZ408r+9OK8UKD/Wlpf1qKQtofUvtDStfq3HHZMkPl5d4/Pp+fp9//",
	"U62W6ec/RSlhw7aHXOo2e17fK7PswZzyk+nUusDNj9suinCAng8rNW9Sx5Fxw4gIT21FDEEQrT+/OUgj",
	"jhv1EZlRRUP2wNNE16bB4ecsgylRRbJEBvyFGoKcOfF5Rt7OKwc5UyhzVzV0yy8eAlpoQYxkKa6wbFHJ",
	"KDC61NzHm6KkOwOLyyBVj5hg8Vr4dXu7coUjPAXhVeFNzW+4q+v7min3Fz5UhP8XuS0G6H44gUxQjLGn",
	"sBLc/bOfTdrRQjmd+3cwq6N4P7n/J8Lg/lWBUv7gIPLD1QCLXIBf2f3g6mEHVBG9LcrUyIFKQUJniYyw",
	"7h+x5DjxHiMphLbP0LToNadKXQuZdoVp2682tq7QS1vx5pezs2MbmWx4chjIUg4Xi1W+ZLm1M/0GsgxV",
	"bE98eslyp5f4etZXYYdYhI7OVC9MnL07RccZcfaaXoCbwS9h3X9w07jv2OISuvxP5tOdYL671viZo2xk",
	"fVum6nP/xXN871TxW2qdRzU/w5iPN2cceOOHYeHXS3AlqSSoXHCFt4LSQlZpGph5YBNZaiHFs7jO98Aq",
	"pirmc/alPdUxlWUF7U8n71z9eLECFVR3u6AKv87IW40JFVZTAPJ7ARjxK+kKNJrx7YV6cM73DRL3tdj3",
	"5uD/hY3/ExvHYNyk45bbtVWt9TveIa7g151sKssa3+2XvN63hnRvWwyeM9wmQRKaZURIkmSC2xfEhlhi",
	"puGCYvdMZ+7+nR5QZvP8OrdCywK2bbkbI77jG+sX3OlSFI4f5TYrUXB93GVs6kyxQnkqp0kPq6KTHaoe",
	"02DSrYemAj2OxLobIJLisrJlSy9hPbWuJWfhMMwEXyP48BoT3YzItM+LLLMhQ8T7IRTB0gBGzl4yHnmM",
	"ED+/Gx6wtHnd4aixM1B6dqJ+O/PFOWAuQBHvALGrVmuul6BZUlX4IKtCWRt+aGrJmNK2WOAVlUwUqvQj",
	"IBhqRg6D2g10bZ0AgmdrfGtAzMkflUtlSjxgN1G7v2a8iIUQuS84vtG9QTvzjH1UBM1UJGMrq5fp2vO1",
	"qGWUCUzupZfgNZggJgwkRlGvhAQUqgi9oixDyxYx7M3SDlNE5PT3Akqn7gXCgQYrfILDv6tQBks733Dg",
	"eaTWF4LampHYmW0lQUsGV/Yu5/BF+4iWEpIK70cWKzYPKxFcMaWBazuWAcs5L519HDzK3ErreYlm3TZp",
	"MSWYb4PyBOWEkjlce9OD3dwcS+VZlPit9x53a2mrp4tZ+xyus9xJi0qvwtjM4sTmuugK015ykfYlIJRs",
	"pqTgGShF1qKw8EhIgJWodKKm0XUoJxBGVXU8TryijDO+eKthdWSYUpsA223KEPWSzlRxocx2m29Icg56",
	"3I7q4WSzKU48caKZ336/wFK7d79aEvJlY1LHmoT0Vk3Po6amU5P6S8g9UIoUNjsQqdei1wzjtwJ1x4Lj",
	"keIpESum3cNZaGQFyWjG/mVfY64BirtrzWbkzy6R9QISaqRAq5aiZ3BZ8Eszkqi+IgocPjFtFBv9pVqP",
	"BIc6S5fNNdmFlGbenVbigwZEZpOZKSdXL2cv/4OkAuE2o1RzWNpnXAM322gWUYrCMUr5HpRmK8zY/N6e",
	"QfYv51tNRGb2D4E4wmCE0kJk5pWAjLRrbGsiRx4hS3s5TXSvh01iWs97LLJ1Pw/XBq711gmrvhl81e8q",
	"I0jmhr/gm1XR+8qeL3euFPZwfNIZO7CtfXcqEk3EudCVpWvHcOOqsX1lZh3GGkdzUP27VmdsBUrTVd6/",
	"qEsKGezYdbHhOZ1DYnlYUvKQWhBOkJgePLVTqpPKCC4upoMcN9/0ssrnjJwATfeMgNDz9Z1bx4H7ovU2",
	"tugS1l6eyQovARilMbjFhVxQbo4ovtlFNSyENP/8s0pEbn+1bPcv5XUc29+4nSLUnF3bmPH1mkNUlg3i",
	"n6gm4hofE8MwNvu7Ed7IOcbz7JupzifEIrnraf7w/u7wFKK04/CH07pSIMw/8Yfc8zsVhL1VZSWraLp+",
	"hpdjI/UGCbTVW2D9tWGRxxXUIP65NFCHwc40TbGYT55ZJUXa0OTPUWtjzDxzSP736ccP5FggJrpt60h8",
	"cRit7KMFoSnKYg6aWUs9QGt0hze9bW0+cW8Q9CsKGAvB9w8T9Cp7hY13Lnf3xMvZtV6N6DxXX2/Ju12K",
	"1w1986JmWIo821p9LRNSXSJD3ewYnOAF0854FD21JxvMmiehGTNIKviZ6dDEaauxoKkLqkc0xvjkMc/g",
	"2ecZVCdoWLJB0O9uMw6qgeNpB/Xv9dyD8hsbM4kePwNBNnaj581YcvsxGeEbTUZo8JxaPGgPn0npbutT",
	"87l341O1rNpugbojtr/ZYliAfyWv9I7yD7rcPia/PtjDZt56efgwA6lPilhgbKOsYFOHWxYryvfKCneN",
	"LBZEnxk7nvJedBlXXntje1hcRVyBDOJ76BVIugBbjApdDT491z+CYCZmfDEjPyEJHHhDTRhu2AginDZD",
	"CKf1AMJpLXxwVo8ePD9P/70zcHA6yUEm5uZadGiz1XeDOrss63SRbLEAqaLotGuyz9pdQZ+CybVNP3Wd",
	"4pUB/YjBXtXWUbcfbaWw2mRBNFv0fQEsxtovSq1zkmrgzibBjJ1tLCjBarz+GEtDWdl3ds2fR8efOo/w",
	"8aeY9dcWjutUrzuKynljdFe/blN1lRnj02achj3sRYKO1Wzj/Zvg2mJo6MDETWSXOgrBepa3ye6AjYgs",
	"sBLpR++ptb/m6E61RIJSkGUqg20RFe+NCF7hbkSfsqSrPGN88daIsFexMo0lK70AfQ3ASxMKdjXrujfu",
	"SN4XCuWwdtD3bIe465q/P8DLNNzLCEo2saXTNU9iAkX1tVl1cA4Sjf5aWK+98wBjzJhN1gsMIFrYeC70",
	"Vzv5F/WcslL6qCqNxpDRGBI+XD/QHBL0vGuDSDW0N4mMp/VxDRuu75ong69Z5PSjaeObNW00OEhnhnB3",
	"jDgtS8jXMkoaOjp5i6/k+BbTc65rOSjVGdWUcRveF7v7bbg9F+dcFRe+OzMnEB8RQFAaY9nQAT8CFqlC",
	"CeScu2Af/wDZk4hTb6dCR1J3XCCEdK3a+B4WXd43g7pBMJ12pWaboZalil/dzk5Ed+N9Gws4eHPJkVit",
	"WEciqI0xwwZkSdWyqoVm4IA0vvN+5J83hM+UowfRMbHB+4RmDTB4narlTilXuWRXVMOvsD6mSuVLSRV0",
	"J0/Z71ZzUsvjsu9TyJmqA7Qtucmtm5ye/tI/v+kmjvgd0zVUuGVbLMn3lKxhVt9wbfvUjR1TNqpFRam0",
	"gyE5JsSsJqoLyZ1cgg+p0MxX7kwF/86/ROEeqg+Cr3pWWexj2624nRV9fMxQRwAVVXEj8oomS8ahc6rr",
	"5boxgcGBuyvO8anzQkL1rIONtmWqCkO3KZ42QBbja+vsuwpePyQnCCZJMipt2JYPYXCLNQeDXBQGy2Aj",
	"dcUVSMlSIExvea4lup0+wK1EHvmI6QAH5HxyWiQJKHU+MWJJsNJ7l/SMWrRHebqn/JMXPQ75mSvP9Dq0",
	"idbyluMlYrYk92xIYepMPuxnOI4CXMI46VhRDdiuRiHIXW2C/LLPAfo6lcpGg7ppKowjJL5Q1uiNH01M",
	"o4mJqv3G0RlmZWp2vltDU2P0ePhNpFE9BqfRYIzDeXRzVWxHeqltzXtgtFp9o1arGFNqFziIl/Q+K582",
	"u14KBeWN78/nHAMGxPZiJXb8PuBVj7X1ym4Kq31Ot/CzXcwr5Yodl7qDWJy7fPza0bp9PahPvtEQS8bn",
	"mxt8I9s+OpmxBLg1SNhEmslhTpMlkFezFxOn1078ybq+vp5R/DwTcrHv+qr9d2+P3nw4fbP3avZittQr",
	"fK5AM52Z4T7mwIndT/K+qlF6ePx2Mp1c+UtlUnD3pKmricRpziYHk7/OXsxeOmMc4tQc0v2rl/u00Mv9",
	"KpNiEaPzn0Hb8iS1kP+wus7b1Cy40MtS2Pb5oTjZqxcvfM402IzV4EHq/f9xKqnd0m0bHsyCG9DIzPvV",
	"rPuHl3+P3K8FGnt1uQqDIxyihosrmrHUFeWNYuM318CixJaRiaHCt0Os+5oeeGKZGWYJNAXpy5baLjav",
	"2CG3QkeTSD/H0ds43ZhZjKtBlLx42dWG8arVbogLXsRw76z4y8eOlkHsrQ37ey2r1DCBo2qwUzuYT69q",
	"Yvk1DtDZXt0nGZYCaBcJWnzfyVz2cY7IVJ+4oUHM9kutOYAuUKnu3BBUcqNkjULsRlzWkW+u4o3NG0Tf",
	"XdKzbGhkUVsExztTikwH0o61r4bVDdydgSOYATBx1la/0M1G3/l0/u9c6rUzXuUSrrBURD2vHZ/GnxxM",
	"EKDqmJZ1HzYd0GksU9Umvrs4FC1Zoqt0dPSsuioEPhXYJqIy6R4EmpHXMKeIEC0IXIFcl+U9YoBmtTIj",
	"g6AN60+Gyfl2O0pAw5IBVTmAs6poA+a221z0bvTXuhM2r+89fGFK20Eb1RgwQHgJvFXesiInDAUKKh0g",
	"hjrxxVaYjFXhKfR7/PVVzO/x+R4ZTOfZQuV0A995cf9850eakuDRt6fM63KhoiUybJ2KAMnEYbnF6Owz",
	"QJtuJTfajyJd3//2W9xUUqqWBdw8Bh120+CrO6SHQdPbrUotDK8eB4bDJIG8BOLvd3cw2s8tRibPJNB0",
	"jdlg0gExcoSQI/SSWvf/MJfCTS/hNcJCyI4C6zahKYwO2TwtXnAYeFHeb66oWZ1x7KBlPBZTeQSSMpP+",
	"cP+TfhD6J1HwW0vw5ug3ihcnvXWpE6DpzoRZ2W2qWhsyQqmtUW9Pp9NJwdnvBby1xiK8DUfSfcKkmxvt",
	"rE28OZXavtZjjXYNQu5vFMCCLHfCYrvXcYcMtq/kuId4+/dh+1YrTnPjBMdRTgzlxGciHT1hflBErzKs",
	"IrQzEzix/e9a0rqH+2sgGxgVyJExPAvGMERP2w/fBe9W2Ph6Z37yGvj6K2AmozD8RGi80/Do3uTfmRDt",
	"Q/Zf0cU2Et5DEp6Z8B/3P6F1fh4JPs9YohvOYPcq6Q6eX/dCaYfZrPr6TJ26FrFbPLhdOHzHlK6+jb7Z",
	"r9U3e0jmLHP7EYXVv3Ts6hjX0Gy7uqrHhSKXsB4Kuu35Ew5Ug7x/Kc7R3byju/luSRdrNg/dflvo+bFu",
	"fcvARge4u+n/+iCiha+q03UXxQVdW72eUHchdXjVy4/3YQFxg/cyd7y8l1lH48LjiKMROm0LqEOcvB1E",
	"HAqmQzSvssdTV7O6iflZera2SeARD2wH5ZwATfvRjbXokJF8viny6fCCosPOP1NR0lAapyFsPJz5pHdO",
	"Pd+MD3M7vY4+gQc8I/09g51cFhs/hQv6ccXbhzsioyg9mpIfTHbfD975iQpkbs/ck5MiQ7MOtxbnCLfA",
	"xv45oG9eLivfPRrFs7705p8J6iS4hTM/zossK5+fsxm0cyH7yXU/g448f7WFHD/cl4Q37SwbZh/mbL6c",
	"FLcbYtuTVtPHIf8IdjfcZz+0d/mDIB6Q8TZ4OrdBVUmkWztXtYJPA/T0U1+EabTyjCoIqiCDSSlQRp4C",
	"NT0XlWTUEB5FdIIyYe0WxQCqrLeusJBWXtwzjhBpoXxLsEiFO7Ixzz+K4zGGZMzvH/P77y6fdwxr6MPM",
	"NufzV31sfaqNwQftjOr7kYo6MrcfLiShV+p4LXd+TFt/PiESsXO2UYwbEjjRljD6inFDdKPoLE9d6+51",
	"Mp6lAj5AjI1EXFR4jVpzBhOaDZzlC5C5ZPZiqdPcSHLfKskN8ED3YHTOAHRHnO6rSELdUfR5FIp/TIlr",
	"NFE9ygnvI+bUkks3xzq7hm2LcOzURvP6njVvOPSIfmweUQdkNCk/qJvx1auHWGUuRQJK0YsM3nDN9Ppu",
	"eMVtPJDbmURUfB3uSRol12cuud6GAuMi7BMjwuctyI4HIGTW+LbCLq7Hn2zHuLmq/PhMPY3uxYqN3sUO",
	"BL5jSpefRifi6EQcs7a/7axtPOyjd7OLgW7Jn0bsdXgw/bf7kHjs2A/sqQwmHW1lj+0Y9CTaEqb2/8D/",
	"3+z755/c80O7SFnNF6S6BK7mS27bZAdzGSDb8zd7a6JZXOOYB2fq8fXepy0FNvZ/izy4favNJfGEN3o6",
	"CqijgDpGuQ3hKbGHVUcpcAMD7X/ZDgnDafLEfpfsrVnv/XHe0JTYc9YnZc9uvS87GvOGSRSRwJ+tRH4C",
	"NP16SPzDSOLPhMQjPL8/a4/bBwIr9RCvjO/w1Gmr004wlvq94wk3Wgb68+Y4lRqG3ItGI8UW7pJUW7yX",
	"8SQrUkDBe7Wicl1PrVde7J+HQDREcZq6zGF1aseIqS8XQmRA+XhcHpABB6bXIVW45lESxraD+ez8rvns",
	"N1OCayupjoFXD3c8+oc/d/F3bPv4Ysijukce7HCMnphRtLsr0a5LJ7lViOMWKXB4FNmor3zDN8xQKqru",
	"midASM/jxnmmhBswRwm5UEwLyXZ68uYk7B434jSaPFNXc4nn9RYvs9yE0XdM6QY+xwjE0cE7OnhvUUvR",
	"n8vRt7uRY20J8wtax2P9TsIG9yFfBBM8cNRfc+ZR4Xzs0L8a7XZIO0OcVBuouyHkrIdI7bVhn7oOuJnK",
	"n6U83UeoiziTNlDTCdB0pKWRloa5djYQlPN9PB2K+mY8Pf1oeLQwP/C56e/z2ciGscPXeG7uT2B+2KMz",
	"CujP4LzWRHMlCpmAWvNkN0uk7X+65kmnkF41edamyArTW42RQdO4MbKG9dEYORojR2PkLe6p6jSN5sgt",
	"XGurQXID6/ImyRrzuh8ZK5jiwc2SzblHuefxDZM1Ku6Sf4bZJjcQelvwGabJ1IZ++lalzQT/TO1KfaS9",
	"qJVyA11ZO+VIVSNV+dt4mL1yA2k5G97Toq1vyGrZj5pHO8iDn6AhlsuNrNnZLr/OE3SfsvVDH6NRmn8m",
	"p9d8sgYQe7wKmU0OJvuTm883/z8AAP//GTuhfYVkAQA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
