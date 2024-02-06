package main

import (
	"cor/common/logic"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
	"io/ioutil"
	"strings"
)

func ExtractText(imageBase64 string) (string, error) {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	credential := common.NewCredential(
		"AKIDynSvCG8nERYTAcqsjNaZLtDJJ6f072iJ",
		"RjphKjxzTr2340mLm16zl5JMBsNlxxjK",
	)

	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ocr.tencentcloudapi.com"

	// 实例化要请求产品的client对象,clientProfile是可选的
	client, err := ocr.NewClient(credential, "ap-guangzhou", cpf)
	if err != nil {
		return "", fmt.Errorf("failed to create ocr client: %v", err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := ocr.NewRecognizeTableOCRRequest()

	// 设置要识别的图片内容
	request.ImageBase64 = common.StringPtr(imageBase64)

	// 发起识别请求
	response, err := client.RecognizeTableOCR(request)
	if err != nil {
		return "", fmt.Errorf("failed to recognize table ocr: %v", err)
	}

	// 提取响应中的表格识别结果
	tableDetections := response.Response.TableDetections

	// 定义一个空字符串变量，用于存储所有单元格的文字内容
	var text strings.Builder
	for _, detection := range tableDetections {
		for _, cell := range detection.Cells {
			// 将每个单元格的文字内容拼接到 text 变量中
			text.WriteString(*cell.Text)
		}
	}

	return text.String(), nil
}

func main() {
	firstKeywords := []string{"取件码", "取货码"}
	secondKeywords := []string{"顺丰快递", "EMS快递", "申通快递", "圆通快递", "韵达快递", "极兔速递", "中通快递", "尼乔卫生站", "邮政快递", "韵达快递", "圆通速递"}
	// 读取本地图片文件
	fileData, err := ioutil.ReadFile("./images/7.jpg")
	if err != nil {
		panic(err)
	}
	// 将图片内容进行 Base64 编码
	imageBase64 := base64.StdEncoding.EncodeToString(fileData)
	text, err := ExtractText(imageBase64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("识别结果:%s\n", text)
	result := logic.IsKeywordBefore(text, firstKeywords, secondKeywords)
	if result {
		contents := logic.GetSecondContentMap(text, secondKeywords)
		jsonData, err := json.Marshal(contents)
		if err != nil {
			fmt.Println("JSON 编码失败:", err)
			return
		}
		fmt.Println(string(jsonData))
	} else {
		contents := logic.GetFirstContentMap(text, secondKeywords)
		jsonData, err := json.Marshal(contents)
		if err != nil {
			fmt.Println("JSON 编码失败:", err)
			return
		}
		fmt.Println(string(jsonData))
	}
	return
}
