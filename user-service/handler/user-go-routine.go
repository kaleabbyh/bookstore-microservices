package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// func SendBulkSmsHandler(gatewaySvc gateway_merchant.IGatewayMerchantService, smsSvc sms_svc.ISmsRecordService) echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		const senderID = "GU_777aea01-36ff-4ba8-b572-df172433ba85"
// 		const pageSize = 1000
// 		message := "ፈጣን እና ቀልጣፋ የሆነውን የሳንቲም ፔይን የክፍያ አማራጭ በመጠቀም ክፍያዎን ይፈጽሙ።\nለሚገጥምዎ ማንኛውም ችግር 9360 በመደወል መፍትሔ ማግኘት ይችላሉ።\nሳንቲምፔይ በጣም ቀላል።"

// 		// Retrieve the last processed phone number from the database if it exists
// 		record, err := smsSvc.GetSMSRecordByPhoneSenderID(senderID)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 				"error":   "Failed to retrieve last processed phone number",
// 				"details": err.Error(),
// 			})
// 		}

// 		// Flags for tracking states
// 		resumed := (record.Reciever != "")

// 		// Start the SMS sending process in a background goroutine
// 		go func() {
// 			pageNumber := record.PageNumber
// 			totalProcessed := 0
// 			resume := resumed

// 			for {
// 				// Fetch a page of gateway users
// 				response, err := gatewaySvc.GetAllGatewayUsers(pageNumber, pageSize)
// 				if err != nil {
// 					fmt.Printf("Error fetching gateway users: %v\n", err)
// 					return
// 				}

// 				fmt.Println("Total phoneNumbers: ", response.Total)
// 				fmt.Println("Total Pages: ", response.LastPage)
// 				fmt.Println("Current page: ", response.Page)

// 				for _, phoneNumber := range response.Data {

// 					// Skip already processed numbers if resuming from the last recorded point
// 					if resume {
// 						if phoneNumber == record.Reciever {
// 							resume = false
// 							continue

// 						} else {
// 							continue
// 						}
// 					}

// 					sms.SendGetSms(message, phoneNumber)
// 					if err := smsSvc.CreateOrUpdateSMSRecord(phoneNumber, message, senderID, pageNumber); err != nil {
// 						fmt.Printf("Failed to send SMS to %s: %v\n", phoneNumber, err)
// 					}

// 					fmt.Printf("Sending bulk message. Page: %d phoneNumber: %s\n", pageNumber, phoneNumber)

// 					totalProcessed++

// 					time.Sleep(200 * time.Millisecond)

// 				}

// 				// Stop if this is the last page
// 				if response.Page >= response.LastPage {
// 					fmt.Printf("Bulk SMS process completed. Total processed: %d\n", totalProcessed)
// 					break
// 				}

// 				// Move to the next page
// 				pageNumber++
// 				time.Sleep(5 * time.Second)
// 			}

// 			fmt.Printf("Bulk SMS process completed. Total processed: %d\n", totalProcessed)
// 		}()

// 		// Respond immediately based on conditions
// 		if resumed {
// 			return c.JSON(http.StatusAccepted, map[string]interface{}{
// 				"message": "Bulk SMS process resumed. It will run in the background.",
// 			})
// 		}

// 		return c.JSON(http.StatusAccepted, map[string]interface{}{
// 			"message": "Bulk SMS process started. It will run in the background.",
// 		})

// 	}
// }

func SendBulksms(msg, url string, phones []string) {

	go func() {
		for _, phone := range phones {

			jsonData, err := json.Marshal(map[string]string{"phone": phone, "message": msg})
			if err != nil {
				continue
			}
			byteBuffer := bytes.NewBuffer(jsonData)

			req, err := http.NewRequest("POST", url, byteBuffer)
			if err != nil {
				continue
			}

			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				continue
			}

		}
	}()

	fmt.Println("Bulk SMS sending started in the background.")
}
