
package bitcoincash

import (  
    "fmt"
    "time"
    "strconv"
    "net/http"
    "encoding/json"
    "log"
    "strings"
    "blocksdk-go/base"
)

type Bitcoincash struct{
    base.Base
}

func getBlockChain(request = map[string]interface{}) *Bitcoincash {

	return Base.request("GET","/bch/block")
}


func getBlock(request = map[string]interface{}) *Bitcoincash {

    value, ok := request['rawtx']
    if !ok {
            request['rawtx'] = false
    }
    value2, ok2 := request['offset']
    if !ok2 {
            request['offset'] = 0
    }

    value3, ok3 := request['limit']

    if !ok2 {
            request['limit'] = 10
    }        
    return Base.request("GET","/bch/block/" + strconv.Itoa(request['block']) + "",{
            "rawtx" : request['rawtx'],
            "offset" : request['offset'],
            "limit" : request['limit']
        })


}


func getMemPool(request = map[string]interface{}) *Bitcoincash {

        value, ok := request['rawtx']
    if !ok {
            request['rawtx'] = false
    }
    value2, ok2 := request['offset']
    if !ok2 {
            request['offset'] = 0
    }

    value3, ok3 := request['limit']

    if !ok2 {
            request['limit'] = 10
    }  

    return Base.request("GET","/bch/mempool",{
            "rawtx" : request['rawtx'],
            "offset" : request['offset'],
            "limit" : request['limit']
        })
}


func getAddressInfo(request = map[string]interface{}) *Bitcoincash {


    valuee, okk := request['reverse']
    if !okk {
            request['reverse'] = true
    }

    value, ok := request['rawtx']
    if !ok {
            request['rawtx'] = false
    }

    value2, ok2 := request['offset']
    if !ok2 {
            request['offset'] = 0
    }

    value3, ok3 := request['limit']

    if !ok2 {
            request['limit'] = 10
    }  

    return self.request("GET","/bch/address/" + strconv.Itoa(request['address']) + "",{
            "reverse" : request['reverse'],
            "rawtx" : request['rawtx'],
            "offset" : request['offset'],
            "limit" : request['limit']
        })

    
}
func getAddressBalance(request = map[string]interface{}) *Bitcoincash {


	return Base.request("GET","/bch/address/" + strconv.Itoa(request['address']) + "/balance")

    
}
func listWallet(request = map[string]interface{}) *Bitcoincash {

    value2, ok2 := request['offset']
    if !ok2 {
            request['offset'] = 0
    }

    value3, ok3 := request['limit']

    if !ok2 {
            request['limit'] = 10
    }  
    return Base.request("GET","/bch/wallet",{
            "offset" : request['offset'],
            "limit" : request['limit']
        })
}
func createWallet(request = map[string]interface{}) *Bitcoincash {

	    value, ok := request['name']
    if !ok {
            request['name'] = nil
    }


    return Base.request("POST","/btc/wallet",{
                 "name" : request['name']
    })
    
}
func loadWallet(request = map[string]interface{}) *Bitcoincash {
	    return Base.request("POST","/btc/wallet/" +strconv.Itoa(request['wallet_id'])+"/load",{
            "seed_wif" : request['seed_wif'],
            "password" : request['password']
        })
    
}
func unLoadWallet(request = map[string]interface{}) *Bitcoincash {



     return Base.request("POST","/btc/wallet/" + strconv.Itoa(request['wallet_id']) + "/unload")
}


func createWalletAddress(request = map[string]interface{}) *Bitcoincash {

    value, ok := request['seed_wif']
    if !ok {
            request['seed_wif'] = nil
    }

    value1, ok1 := request['password']
    if !ok1 {
            request['password'] = nil
    }

        
    return Base.request("POST","/btc/wallet/" + strconv.Itoa(request['wallet_id']) + "/address",{
            "seed_wif" : request['seed_wif'],
            "password" : request['password']
        })
}
 

func getWalletBalance(request = map[string]interface{}) *Bitcoincash {

    return Base.request("GET","/btc/wallet/" + strconv.Itoa(request['wallet_id']) + "/balance")      

    
}
func getWalletTransaction(request = map[string]interface{}) *Bitcoincash {


    value, ok := request['order']
    if !ok {
            request['order'] = 'desc'
    }

    value1, ok1 := request['offset']
    if !ok1 {
            request['offset'] = 0
    }
    value2, ok2 := request['limit']
    if !ok2 {
            request['limit'] = 10
    }
    value3, ok3 := request['category']
    if !ok3 {
            request['category'] = 'all'
    }


    return Base.request("GET","/btc/wallet/" + strconv.Itoa(request['wallet_id']) + "/transaction",{
            "category" : request['category'],
            "order" : request['order'],
            "offset" : request['offset'],
            "limit" : request['limit']
        })
    
}
func listWalletAddress(request = map[string]interface{}) *Bitcoincash {
	    value, ok := request['address']
    if !ok {
            request['address'] = nil
    }

    value1, ok1 := request['hdkeypath']
    if !ok1 {
            request['hdkeypath'] = nil
    }
    value2, ok2 := request['offset']
    if !ok2 {
            request['offset'] = 0
    }
    value3, ok3 := request['limit']
    if !ok3 {
            request['limit'] = 10
    }


        
    return Base.request("GET","/btc/wallet/" + strconv.Itoa(request['wallet_id']) + "/address",{
            "address" : request['address'],
            "hdkeypath" : request['hdkeypath'],
            "offset" : request['offset'],
            "limit" : request['limit']
        })
    
}
func sendToAddress(request = map[string]interface{}) *Bitcoincash {
	    value, ok := request['kbfee']
    if !ok {

        blockChain = *getBlockChain()
        request['kbfee'] = blockChain['medium_fee_per_kb']
    }

    value1, ok1 := request['seed_wif']
    if !ok1 {
            request['seed_wif'] = nil
    }
    value2, ok2 := request['password']
    if !ok2 {
            request['password'] = nil
    }



    return Base.request("POST","/btc/wallet/" + strconv.Itoa(request['wallet_id']) + "/sendtoaddress",{
            "address" : request['address'],
            "amount" : request['amount'],
            "seed_wif" : request['seed_wif'],
            "password" : request['password'],
            "kbfee" : request['kbfee']
        })
    
}
func sendMany(request = map[string]interface{}) *Bitcoincash {

    value, ok := request['seed_wif']
    if !ok {

        request['seed_wif'] = nil
    }



    value1, ok1 := request['password']
    if !ok1 {
            request['password'] = nil
    }
            

        
    return Base.request("POST","/btc/wallet/" + strconv.Itoa(request['wallet_id']) + "/sendmany",{
            "to" : request['to'],
            "seed_wif" : request['seed_wif'],
            "password" : request['password']
        })
    
}

func getTransaction(request = map[string]interface{}) *Bitcoincash {

        return Base.request("GET","/btc/transaction/" + strconv.Itoa(request['hash']) + "")
    
}







