//
// This file was generated via github.com/skx/implant/
//
// Local edits will be lost.
//
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

//
// EmbeddedResource is the structure which is used to record details of
// each embedded resource in your binary.
//
// The resource contains the (original) filename, relative to the input
// directory `implant` was generated with, along with the original size
// and the compressed/encoded data.
//
type EmbeddedResource struct {
	Filename string
	Contents string
	Length   int
}

//
// RESOURCES is a map containing all embedded resources. The map key is the
// file name.
//
// It is exposed to callers via the `getResources()` function.
//
var RESOURCES = map[string]EmbeddedResource{

	"data/archive.tmpl": {
		Filename: "data/archive.tmpl",
		Contents: "H4sIAAAAAAAC/6xVTW/kNgy951dwhWx66NpKChQoEtvANmkvRdtFs5c9FbRN20JkSZA4kwwG898L+SueD6CH9iTJJJ9I8T06+/D05+PXb19+gY57XVxlcQGNps0FGVFcAWQdYR03ABkr1lT8rG0Ln33VqS1lcvw22ntihKpDH4hzseEm+UmsTQZ7ysVW0auzngVU1jAZzsWrqrnLa9qqipLh8AmUUaxQJ6FCTfldejtDaWVewJPOBWomb5BJAO8c5QKd06pCVtZIH8L3b70W0HlqcrHfezQtwbX6BNcM9zmkh8N+rxq4VhB3pAPFNf3iqVFvwydTL4syNb2lPgQBQ8m5+Ov5GRqiGhrr4b/CT8Xt90y908gEokZGqUwlqxBS7p0Wh0NsiJw7kpW23k2PUqstVBpDyEU0k58AATL8n15AzBdo21pRPDNtCX6j3n0XIHIiCw5NAQk8UanQwA327gF+9UTwbBt+RU+ZHHwyiUt2Z4knXrUdL+mvCpBY2g1LUXyO6wpk8JlAsGK1JbGEjDyNQTNj8RI0YxukKL5iG448LjbEHzVkAJK12k6tWG1dcWPK4B4y6WYJYTnLJZ58kXENI/3F3e1HAarOReSaYVFkkutVslwP1kk1qxeKxu7uvcLu7tgW3zjwLpJ2uOr+x9uPx+EzN9Kloil0o+dIrQInw/7eWEMPxwCjMIusLPb79BuhPxwyWZ66jIAzjSJg6+3GiXO/94yekPEkq+W+c6hEMfWiWHFmIsCSVtz9bg13h4MUxXz4A3uKKWMBA49n5BLrlga3R7sxHF1GmmdSqwtpx8QHtZw+jtzo4jxm+Hz1L+ErRo3HU1ZM/Plh5k9QNZXo1xr6kCTQYfUCSbIm94WJsLLOowHm2XC1zvOSMKgiw387G3hSyAmjTqs7Pa9ryyT7WVIr2VwQ1WqGNNbyevi5hSLUsCgePSFTDeXuXfsdswv3UraKu02ZVraX4eVNkuuoJ6+CKJZt5Ee63LuGn4bWTWXd7uEcO8Rhmb5Q79JGyfXwjJDvlcydzuQ42zM5/pj/CQAA//8ZL8U/qQcAAA==",
		Length:   1961,
	},

	"data/archive_page.tmpl": {
		Filename: "data/archive_page.tmpl",
		Contents: "H4sIAAAAAAAC/4RUwW7bOBC95ytmCWz2sLUY91Q0lIA0SS9F0aDOJaeCEkcSEYokyIljQ9C/F7QkR6kT9KSh+Thv5s0bi39uflzfP9zdQkudKc5E+oCRtskZWlacAYgWpUoBgCBNBosvxjVwFapWbxFW0PfZd2epHYYUPqAMwyD4CB2fdUgSqlaGiJSzJ6pXn9jyysoOc7bV+OxdIAaVs4SWcvasFbW5wq2ucHU4fABtNWlpVrGSBvN1djGnMto+QkCTM2kIg5WEDGjvMWfSe6MrSdpZHmL8f9cZBm3AOmd9n90FrPVuGLRVuMtCjAwO1efs52YDNaKC2gVYICfKvifsvJGEwJQkybWteBVjRp03bBiSenyWT5RO7adSld5CZWSMOUvXGKaEAEKe1sVmrHGNY8WGcIvwDTv/X4Q0CxG9tAWs4AZLLS2cy85fwteACBtX07MMKPgBI7g8Ep3UsAq6aelYyaIWLkv3RJwVV+m7SPIKMxoiocboHRzJJnJW3MsmvkK8KWZ4JeYhEVd6O8m4CH1xbsvoLwX3s1dlORswnUIhSMFoKLa++JeBVjlLI7fECsFJLYoldbidfLiQJK3Duri1FDRG8C4SKtD2nR1o18uXfR+kbRCy6fmxpXe7L41rfiWSUQPI/niCVi1+Oe1h6vbj3G3UCksZWPE34oAVWjpQxze4l0yCU5jHsZD8jYEsDFc7R0vT+6PDsSZWXAeUSdhy/+KblsjHz5w3mtqnMqtcx+PjjqNvscOgIyuOYbJVduRdpp8cfl45v788zR3TZmWP2Pms1ny5aSnlSyez7wQfd1rw8d/zdwAAAP//RJhDI04FAAA=",
		Length:   1358,
	},

	"data/entry.tmpl": {
		Filename: "data/entry.tmpl",
		Contents: "H4sIAAAAAAAC/5xUwW7bOhC85yv2EXjpobWY9FQ0lIA2SS89NKhz6clYiSuJCEUS5MaxYfjfC1mSo8AJAvQk0hzODmdnrf67+XV9/+fuFlrubHGm+g9YdE0uyIniDEC1hLpfACg2bKnY7bJbx3Gb3ffb/V7J4fcB0xEjVC3GRJyLR64XX8T8yGFHuVgbego+soDKOybHuXgymttc09pUtDhsPoFxhg3aRarQUn6ZXUxU1rgHiGRzgZYpOmQSwNtAucAQrKmQjXcypvRx01kBbaQ6F7tddhepNpv93jhNmyymJOCgPhe/l0uoiTTUPsIMOZbc7Zi6YJEJhEZGaVwlq5Qy7oIV+31vlZy8UqXX21GqNmuoLKaUi/6Y4kgIoPBUl5iw1jdeFEumNcFP6sKHBN+tb1QK6ApYwA2VBh2cYxeu4EckgqWv+QkjKXnAKInHQicaFtE0LR+VzLRILP0jS1F8678zkheYWLVmTT1qWL2BY2ySFMU9NukF4lUz4wszD0RSm/Vo42wZinNXpnClZJiCieUUwH4XC8UahkCJy4v/BRidi77ljkWhJOuZWNaH0zGHM0vekFla36yCTzyIhWEYjprfzorvOnKcVt6t/pUDtV6NPKvax266Prt5+rrRh8+TD8loKjG+/9JIVV+o15neqaQkx6lRs2a80qpZFGvveT4O4Zh9qlkU15GQSUO5fU5UyxzSVykbw+1jmVW+k+lhIym01FE0SRTHZR+47Fh3Tj9m/7zyYXt1yp36mcseqAtZbeR8BnvK55dMiVRymHYlhz/RvwEAAP//CbGSNlUFAAA=",
		Length:   1365,
	},

	"data/inc/add_comment_form.tmpl": {
		Filename: "data/inc/add_comment_form.tmpl",
		Contents: "H4sIAAAAAAAC/5yTXWvdMAyG7/MrhAa7Wk/aju7TMZTSu7KWbVB2acfKialjB1tpm4X89+GTZNBBO+hN8IfeR3ojeZpsA7tzYy5C15FnmOdCtKfy3BgYwxChXs5F2Z7KQjQhdqBqtsFXWNZ7e6StL9eYtKv3FsGaCusciOBVR383HXEbTIU31z9+Iqi6pp6P6lbFRFzhwM3RJ5SFsL4fGHjsqcLWGkN+41iDcK/cQBVO09X17eV32F16juPuyvq7eUYoX9THoAMv9a3LlfY/YRND4nFRbutV+vbN54/vT77mwllpR7IAEBylYCOFlt9UR6LU8oso2eS7fGvkk1xMj4yQ7G+q8Oxsy5m/ua6DUJQc/yFfdsq616Gz8iX0Lelk+ZV1O+vvnoNDHVzqla/wFHOiX3nALrYB01LomDsBsOSjR1aR1ArWwYwIMTykCk+O8cCq8MMx5ixr6DOWDodsQDm79xVGu28Zn5pJg+4sbya23drkm5B4qxOfWBPl2nVR5iGXhegXVwdCSjZ4eLDOgSawex8iGbANKD9CaIBbgsaSMwlUJHDUMGin/N070AMv749yt0AZEymlhSW09HRPMf+yDDY29U6NZHai7GUxTeTNPBd/AgAA///objF82wMAAA==",
		Length:   987,
	},

	"data/inc/blog_post.tmpl": {
		Filename: "data/inc/blog_post.tmpl",
		Contents: "H4sIAAAAAAAC/4STQW/aQBCF7/kVI5f2FGyaozGWUkBVVQRRg9Qj2tiDvcoy63qnadF2/nu1JiYpKsQXez1vvzd+O85K/ZRfAWT1DRRGOTeJflKJrdGEZZRnCuoWt5PI+8Xq+/wbxAtNjyIRON4bnESMv3lYYmFbxdpSSpZwDIU1tk0fjCoex1HufbzWbFAkS1SeJfVN59j8w1BGV5S2uqp5DL90yXX6cTR6H3Z/Wqw+b+5W9+vN7HY9h3imuGM1HabUT33nSNzuh4UlRuLOd3p4FrmCi46vDYMUMtcoOuEqxsq2+2cFeA8Dg1RxDekEDBLEa1U5EAklvQX8cRSM4NBCuJYWWFXuCEHj8FU5MFLwvlVUIQz0NQw4GHRwEe/1FgZa5Bq8RypFjmeUBGzi/fx+ens33yy+LL9CLBKCOLwKq3AEzxv7Dl4WWRI++z8J7JDV0GET5X/Oavr0dzskdhdSmh4kM8XqrbCWFnrgmbzOTei7l0a8P7UVOWJDHheD6Ocs6f6U/tbkH+jBNeOu+jcAAP//eb+BcEUDAAA=",
		Length:   837,
	},

	"data/inc/comments_on_blog_post.tmpl": {
		Filename: "data/inc/comments_on_blog_post.tmpl",
		Contents: "H4sIAAAAAAAC/3RRT2vCMBS/91M8wq5rh7uNtODUw0Dn2IQdJTbPJixNWBIVeeS7D631z9h66Uve7x+/EMGdQdtEBU8lGLSQj1zboo1jEQWkBESg14DfZ9wDpEQEaAJCShlXA9CyZHVHCwysaPHqXJ0EAzgLUekAaKPf80INqowIrUwpI/LCNnjjfhCXegu1ESGcBVmVAfxxr1BI9Mct8ChWBk+jr3iUsNMyqpI9DjpE/3HdNiBMLJmunWVXMFCoGxW7Ofi6ZET5S+1sSqxjbAL6+yOt4kWUZ10eZUWUDzdROX9scDSfzSavi+Xb/GOxHA8XE8jHImJKfOWLmzzABSiP64PbdP45eYd8qu3XwdOjKZl1a2eM27Hq154X4hKgT8OL6Puhb4QXUm//6XDl5L7rhyh/dnKf0oVx+vUv9hMAAP//NnRURzoCAAA=",
		Length:   570,
	},

	"data/inc/css.tmpl": {
		Filename: "data/inc/css.tmpl",
		Contents: "H4sIAAAAAAAC/3RU3a6bPBC85ylWOvqkrxek5LSRKnNU9VUWdgPWMV5qO39FeffKGAj5qXIT1jvj9czYHz5cDP/MNi0jsYMB5Mhub+SkoNVEbEuosP5snBws5bUYcQre9tv4K+G6ABEG2BvBoMDwPpQwtVYG688SAp9DjkY3VkHNNrAroUcibRsF2/f+PLUQ1+IwaLEKrFguYS825F7/YQXbH7HPaMt5y7ppg4L3XSxV4ohd7pD0wSv4Hmur0TZGGonz3agSbiycJqpKDL2A+R7tA7Z4wlpxHZo7tGqjkjC8ko+IHgVa74t10Ed+CSWhhl1lDrwQnFoduFzwuYsj3dwYPyN/IBhWmo+HOLILukYzWxOkH1uxMnH/k6bQqm1R/Ldo7HusVwRTtRZjsPeswHOPDgNHmspI/fn7IGE8S+qM4VBQLO5PhV2RPKPVkI9LYbbh3rG5O0ivdpPze5Ewij/l5Nt62Ni47c/gxWha6R+3g2Ed1ZTkJGT6f802o6CzvJO6K8yi9yblnOmOcg7/NdscLLGLaSYYXihdSQjSKbiN6rU5Tthauo5tWHRdtzWOLyV06BptZ5ZtVOaGS1EZ/r1V4rgBKqHLc3yu2a+OSSP42jFbQEvwf4fnPB0HdkXRn7/AkMGLZyJdb9K+N3iJ90AeH4pZcniI9h3BNYM3r4krjH4vdLdFbSkp9bR2zT6+Tu/f3wAAAP//e3izNAgFAAA=",
		Length:   1288,
	},

	"data/inc/recent_posts.tmpl": {
		Filename: "data/inc/recent_posts.tmpl",
		Contents: "H4sIAAAAAAAC/4SOsUrGMBRG9zxFyG4y/JveXhCbrdhSA44l1asNhhSadJBL3l0a3d0OH5zDB9sNZ3qjVOS055LBbDcUcEaZy3ekTsWQy13j+7QnelAomA+fPknqX7F5tQqIAWFF5tk+2We3TOOLW/pHZ6XufaFawaz/pVvDy+2gj04xD+OrnaUeQvqqVSGzdqHElvIIJgYUYM74h8yU3q8j1yR+AgAA///+5Bfz2gAAAA==",
		Length:   218,
	},

	"data/inc/rss.tmpl": {
		Filename: "data/inc/rss.tmpl",
		Contents: "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",
		Length:   2126,
	},

	"data/index.rss": {
		Filename: "data/index.rss",
		Contents: "H4sIAAAAAAAC/3SR0WvbMBDG3/VXCBf2Fp2dwlg8xWVrEhgrtMSDPavWORWVpUxS6gzj/33Icd2k2x519/3uu/vEb46Npi/ovLJmmWQsTW4Kwp2s8+1qQ+ix0cbnTtbL5CmEfQ7Qti1rr5l1O8gWiwWkc5jPZ07WM//bBHGcGX+VvIKymrj9wemBkhWgxgZN8JCxDCZxbcWbzVBilW0gliE9F1bWBDTh79HOx5EpNFYeNHoYhRP5fyIhBeHVkzAGNY3Xi0d7CMuk69iDw1od+z4pCA8qaCzKgC9Iv2Oz/3B1vfjs6VdtdxxOTcK1Ms/FGchhqBAu0VdO7YOypljhoxKGCiPpxiHS0tahFQ45nKsIVwEbXxA6fEmJvwrSdU6YHVK2NsEp9H1P6Kmt1bC6Q28PrsK4/d39z/WWsjtlnvs+gUijkRHhME3kMLpwGCP4p8uwy2U4l+OngLpuXd5+eVhT9iO+YwTvwrkAp4DG/8rRVFaiPJtze+pE7XsR4bLKpQjR9lt5/+ljmlG2EmGwfW2NN77dfzp/u9oU5E8AAAD//8Gx+ecEAwAA",
		Length:   772,
	},

	"data/index.tmpl": {
		Filename: "data/index.tmpl",
		Contents: "H4sIAAAAAAAC/3RUwW7bOBC95ytmCWz2sGsx2VPRUALaJL300KDOpadiJI4kIhRJkBPHhuF/L2RLDhM7Jw49z28eHx+l/rr7cfv46+Eeeh5sdaHGBSy6rhTkRHUBoHpCPRYAig1bqpZMK4LvNIR/Eny1vlPy0DiABmKEpseYiEvxzO3ik8hbDgcqxcrQS/CRBTTeMTkuxYvR3JeaVqahxX7zHxhn2KBdpAYtldfF1UxljXuCSLYUaJmiQyYBvAlUCgzBmgbZeCdjSv+uByugj9SWYrstHiK1Zr3bGadpXcSUBOzVl+LncgktkYbWR8iQ08jtlmkIFplAaGSUxjWySangIVix241eydksVXu9maRqs4LGYkqlGNsUJ0IAhae6xIy1vvPijNkpoKtgAXdUG3RwiUO4gW+RCJa+5ReMpOQeoyQeB51oWETT9XxUkmmRWPtnlqL6Mq4ZyRtMbHqzohF1qD7AMXZJiuoRu/QGcdbM+MbMPZHUZjXZmJWhunR1CjdKhjmZWM8BHHexUqzhEChxffW3AKNLMV65Y1EpyToTy3rfnXKYWTLKjOg6guLecTSUjsI+PENtffc7+MSHk0Dx7i/kdPbLqZJJ8/+z5mQ01RjfqTpnHjXkeD86nZmdT1KS42xqZtwZW7PYtN5zHt1wzCm1LKrbSMikod683n7PHNJnKTvD/XNdNH6Q6WktKfQ0UDRJVMdyDEdxnJvTTzm9bHzY3Jxyp/F9FE80hKI1Mn8vI+XrSeb0KHl4mUoevnh/AgAA//8mAQuNAgUAAA==",
		Length:   1282,
	},

	"data/tag_page.tmpl": {
		Filename: "data/tag_page.tmpl",
		Contents: "H4sIAAAAAAAC/4RUsW7bMBDd8xVXAk2H1mLcqWgoAWniLh0a1F46FZR4kohQJEFeHBuC/72QJTlMnaCTjuTju7vHdxLv7n7ebn7fr6ClzhQXYviAkbbJGVpWXACIFqUaAgBBmgwWK0tBY4SNbBpU0Per9e3N/QqyjWwOB8FH1HijQ5JQtTJEpJw9Ur34wtIjKzvM2Vbjk3eBGFTOElrK2ZNW1OYKt7rCxXHxCbTVpKVZxEoazJfZ1UxltH2AgCZn0hAGKwkZ0N5jzqT3RleStLM8xPhx1xkGbcA6Z32f3Qes9e5w0FbhLgsxMjhWn7Nf6zXUiApqFyBBTin7nrDzRhICU5Ik17biVYwZdd6ww2EQjs/KidKp/VSq0luojIwxZ8MxhokQQMjzutiMNa5xrFgTbhF+YOc/RPhmXCOil7aABdxhqaWFS9n5a/geEGHtanqSAQU/YgSXp0RnNSyCblo6VZLUwmXpHomz4mb4JiQvMKFq9RYH1Bi9gSPZRM6KjWziC8SrYoYXYh6JuNLbScYk9MWlLaO/FtzPNpXlbMBhFQpBCkZDseXVewZa5Wx4ckusEJxUUiyp4+nkw0SSYRKWJ+/T6H1ROYXF2QQcdwVvl+n1vg/SNgjZxHHq600JSuOaP95FGoWA7J8raFWyc97I1PLnueWoFZYysOJ/iQNWaOmYOr6SO80kOIX5TRLdX3mVxHW1c5Q6359sjjWx4jagJFRQ7p/N0xL5+JXzRlP7WGaV63h82HH0LXYYdGTFKRy8lZ3ypvSTzS8r5/fX59xxGK/sATuf1Zqn4zZQPncym0/wcbAFH/+efwMAAP//1r3kj04FAAA=",
		Length:   1358,
	},

	"data/tags.tmpl": {
		Filename: "data/tags.tmpl",
		Contents: "H4sIAAAAAAAC/3SUwW7jNhCG7/sUUwJND63FdU/FhhKw9aZAsUUb1OmhpwUljiQiFEmQY8eOoHdfUJZsJU5OpjzDmV//fCPxw5d/Ng//399BS50pPoj0A0baJmdoWfEBQLQoVToACNJksHiQDWyM2ynBT3+cgh2ShKqVISLlbEf16je2DFnZYc72Gp+8C8SgcpbQUs6etKI2V7jXFa7Gh19AW01amlWspMF8nX2cSxltHyGgyZk0hMFKQgZ09Jgz6b3RlSTtLA8x/nzoDIM2YJ2zvs/uA9b6MAzaKjxkIUYGo/qc/bvdQo2ooHYBFplTy74n7LyRhMCUJMm1rXgVY0adN2wYkkd8NkmUTh0nqUrvoTIyxpylMIapIICQ17rYnGtc41ixJdwjfMXO/xThd+MaEb20BazgC5ZaWriRnb+FPwIibF1NTzKg4GOO4PLc6ErDKuimpbOShRYuS7cjzorP6XdR5EVOqFq9x5R1Or3Om5rJivQeZ/s5ySZylsiJLy686W144e1Ylyu9n1xdHH1xY8vobwX3M6CynHlMT6EQpODEF1t//JGBVjlLBFhiheCkFtpJjdEJy4tDol0vkW/X50jfB2kbhCy911ksjJOCSMcEV+0sraJ+xk8JrYetfsZh8AdWXDwdzen7u+3m8/3dt7/+/PvrWDEhMRH6X0SVrm/cztIwAOkOIyvmO1N6cnZC4KIQrboog+tXnsz5dTYnaoWlDAtA3psSVmjpm3eRpnFB9k4nwSnM01tM6I35LXCtnaPlyvjzfmBNrNgElIQKyuOFzpbIx0+cN5raXZlVruPx8cDRt9hh0JEV52PyKjv3XZaf9uOmcv54e107pr3MHrHzWa35ck9P9vtXmAp++iIIfvrCfg8AAP//z5YE4XIFAAA=",
		Length:   1394,
	},
}

//
// Return the contents of a resource.
//
func getResource(path string) ([]byte, error) {
	if entry, ok := RESOURCES[path]; ok {
		var raw bytes.Buffer
		var err error

		// Decode the data.
		in, err := base64.StdEncoding.DecodeString(entry.Contents)
		if err != nil {
			return nil, err
		}

		// Gunzip the data to the client
		gr, err := gzip.NewReader(bytes.NewBuffer(in))
		if err != nil {
			return nil, err
		}
		defer gr.Close()
		data, err := ioutil.ReadAll(gr)
		if err != nil {
			return nil, err
		}
		_, err = raw.Write(data)
		if err != nil {
			return nil, err
		}

		// Return it.
		return raw.Bytes(), nil
	}
	return nil, fmt.Errorf("failed to find resource '%s'", path)
}

//
// Return the available resources in a slice.
//
func getResources() []EmbeddedResource {
	i := 0
	ret := make([]EmbeddedResource, len(RESOURCES))
	for _, v := range RESOURCES {
		ret[i] = v
		i++
	}
	return ret
}
