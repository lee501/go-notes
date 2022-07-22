package HttpResetResponseBody

/*
   //perform http request
   resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewBuffer(requestData))
   defer resp.Body.Close()
   utils.CheckErr(err)

   // read the response body to a variable
   bodyBytes, _ := ioutil.ReadAll(resp.Body)
   bodyString := string(bodyBytes)
   //print raw response body for debugging purposes
   fmt.Println("\n\n", bodyString, "\n\n")

   //reset the response body to the original unread state
   resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))


   // Step 3
   oR := new(jsonResponse)
   json.NewDecoder(resp.Body).Decode(oR)

*/
