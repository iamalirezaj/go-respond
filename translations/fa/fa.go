package fa

var Messages = map[string]interface{}{

	"success": "موفق",
	"failed":  "نا موفق",

	"errors": map[string]map[string]interface{}{
		"success": {
			"insert": ".درخواست با موفقیت درج شده است",
			"delete": ".درخواست با موفقیت پاک شده است",
			"update": ".درخواست با موفقیت ویرایش  شده است",
		},
		"failed": {
			"insert": ".درخواست با موفقیت درج نشد",
			"delete": ".درخواست با موفقیت پاک نشد",
			"update": ".درخواست با موفقیت ویرایش  نشد",
		},
		"1001": {
			"message": ".درخواست مورد نظر پیدا نشده است",
			"type":    "error",
		},
		"1002": {
			"message": ".کاربر مورد نظر موجود نیست",
			"type":    "error",
		},
		"1003": {
			"message": ".نوع کاربری وارد نشده است",
			"type":    "error",
		},
		"1004": {
			"message": ".ورودی مورد نظر تکراری است",
			"type":    "error",
		},
		"1005": {
			"message": ".نقش کاربر مورد نظر تکراری است",
			"type":    "error",
		},
		"3001": {
			"message": ".شما به سیستم وارد نشده اید",
			"type":    "error",
			"cat":     "auth",
			"short":   "not-logged-on",
		},
		"3002": {
			"message": ".نشان شناسایی شما نامعتبر است",
			"type":    "error",
			"cat":     "auth",
		},
		"3003": {
			"message": ".نشان شناسایی شما نامعتبر است",
			"type":    "error",
			"cat":     "auth",
		},
		"3005": {
			"message": ".نشان شناسایی شما نامعتبر است",
			"type":    "error",
			"cat":     "auth",
		},
		"3006": {
			"message": ".نشان شناسایی شما نامعتبر است",
			"type":    "error",
			"cat":     "auth",
		},
		"3007": {
			"message": ".نشان شناسایی شما نمایش داده نمیشود",
			"type":    "error",
			"cat":     "auth",
		},
		"3008": {
			"message": ".نشان شناسایی شما نامعتبر است",
			"type":    "error",
			"cat":     "auth",
		},
		"3009": {
			"message": ".نشان شناسایی شما نامعتبر است",
			"type":    "error",
			"cat":     "auth",
		},
		"3010": {
			"message": ".زمان استفاده از نشان شناسایی شما گذشته است",
			"type":    "error",
			"cat":     "auth",
		},
		"3011": {
			"message": ".نشان شناسایی شما نامعتبر است",
			"type":    "error",
			"cat":     "auth",
		},
		"3012": {
			"message": ".نشان شناسایی شما نامعتبر است",
			"type":    "error",
			"cat":     "auth",
		},
		"3013": {
			"message": ".Payload معتبر نیست",
			"type":    "error",
			"cat":     "auth",
		},
		"3014": {
			"message": ".Claim معتبر نیست",
			"type":    "error",
			"cat":     "auth",
		},
		"3015": {
			"message": ".نشان شناسایی شما نامعتبر است",
			"type":    "error",
			"cat":     "auth",
		},
		"5401": {
			"message": ".شناسایی کاربر نامعتبر است",
			"type":    "error",
		},
		"5404": {
			"message": ".صفحه درخواست شده پیدا نمیشود",
			"type":    "error",
		},
		"5405": {
			"message": ".شما به درخواستی که داده اید دسترسی ندارید",
			"type":    "error",
		},
		"5406": {
			"message": ".پارامترهایی که شما وارد کرده اید نا معتبر است",
			"type":    "error",
		},
		"5420": {
			"message": ".خطای اعتبار سنجی",
			"type":    "error",
		},
		"5422": {
			"message": ".نشان شناسایی شما نامعتبر است",
			"type":    "error",
		},
		"5445": {
			"message": ".ارتباط با پایگاه داده مشکل دارد",
			"type":    "error",
		},
		"5447": {
			"message": ".عملیات پاک کردن درست اجرا نشده است",
			"type":    "error",
		},
		"5448": {
			"message": ".عملیات درج درست اجرا نشده است",
			"type":    "error",
		},
		"5449": {
			"message": ".عملیات ویرایش درست اجرا نشده است",
			"type":    "error",
		},
	},
}
