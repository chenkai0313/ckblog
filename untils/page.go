package untils

import (
	"strconv"
)

func Pagination(Counts, pageNow, pageSize int, params map[string]string) string {
	var paramsUrl string
	for k, v := range params {
		if len(v)>0 && v != "" {
			paramsUrl += "&" + k + "=" + v
		}
	}
	var pageCount int
	if Counts%pageSize == 0 {
		pageCount = Counts / pageSize
	} else {
		pageCount = (Counts / pageSize) + 1
	}
	var str string
	str += "<ul class=\"pagination\" style=\"color: #00a299;width: 100%;\">"
	str += "<li><a href=\"?pageNow=1" + paramsUrl + "\">首页</a></li>"
	if pageNow == 1 {
		str += "<li><a href=\"#\">&laquo;</a></li>"
	} else {
		PageNowFront := strconv.Itoa(pageNow - 1)
		str += "<li><a href=\"?pageNow=" + PageNowFront + paramsUrl + "\">&laquo;</a></li>"
	}
	PageNow_string := strconv.Itoa(pageNow)
	//总页小于10页
	if pageCount <= 10 {
		for i := 1; i <= pageCount; i++ {
			if i == pageNow {
				str += "<li class=\"active\" ><a href=\"#\" style=\"background:#00a299;color: white;\">" + PageNow_string + "</a></li>"
			} else {
				i_string := strconv.Itoa(i)
				str += "<li><a href=\"?pageNow=" + i_string + paramsUrl + "\">" + i_string + "</a></li>"
			}
		}
	} else {
		//总页大于10页
		//当前页小于10
		if pageNow < 10 {
			for i := 1; i <= 10; i++ {
				if i == pageNow {
					str += "<li class=\"active\"><a href=\"#\" style=\"background:#00a299;color: white;\">" + PageNow_string + "</a></li>"
				} else {
					i_string := strconv.Itoa(i)
					str += "<li><a href=\"?pageNow=" + i_string + paramsUrl + "\">" + i_string + "</a></li>"
				}
			}
		} else {
			j := pageNow
			page_start := pageNow - 5
			page_end := pageNow + 5

			if pageCount <= page_end {
				page_end = pageCount
			}

			for j = page_start; j <= page_end; j++ {
				j_string := strconv.Itoa(j)
				if j == pageNow {
					str += "<li class=\"active\" ><a href=\"#\" style=\"background:#00a299;color: white;\">" + PageNow_string + "</a></li>"
				} else {
					str += "<li><a href=\"?pageNow=" + j_string + paramsUrl + "\">" + j_string + "</a></li>"
				}
			}
		}
	}
	if pageNow == pageCount {
		str += "<li><a href=\"#\">&raquo;</a></li>"
	} else {
		PageNowB := strconv.Itoa(pageNow + 1)
		str += "<li><a href=\"?pageNow=" + PageNowB + paramsUrl + "\">&raquo;</a></li>"
	}

	PageCount_string := strconv.Itoa(pageCount)
	str += "<li><a href=\"?pageNow=" + PageCount_string + paramsUrl + "\">尾页</a></li>"
	str+="	<li><div class=\"summary\" style=\"display: inline-block;float: right;\">共<b>"+strconv.Itoa(Counts)+"</b>条数据</div></li></ul>"
	return str
}
