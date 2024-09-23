#!/usr/bin/python3

import json

ENDPOINTS_METHODS = {
    # endpoint url : (golang method name, response type name, result type)
    ### address
    "/address/{id}": ("GetAddress", "resultGetAddress", "*AddressInfo"),
    #"/address/{id}/txs": ("GetAddressTxs", "resultGetAddressTxs", "[]TxInfo"),
    #"/address/{address}/stakes": ("GetAddressStakes", "resultGetAddressStakes", "[]ValidatorStake"),
    #"/address/{address}/rewards": ("GetAddressRewards", "resultGetAddressRewards", "[]Reward"),
    ### nft
    #"/nfts": ("GetAllNFT", "resultGetAllNFT", "[]NFT"),
    "/nft/{collection}": ("GetNFTCollection", "resultGetNFTCollection", "interface{}"),
    "/nft/{collection}/{id}/txs": ("GetNFTTransactions", "resultGetNFTTransactions", "[]TxInfo"),
    ### tx info
    "/tx/{hash}": ("GetTxByHash", "resultGetTx", "*TxInfo"),
    "/txs": ("GetTxs", "resultGetTxs", "[]TxInfo"),
    ### coins
    "/coins": ("GetCoins", "resultGetCoins", "[]CoinInfo"),
    "/coin/{coin}": ("GetCoin", "resultGetCoin", "*CoinInfo"),
    ### blocks
    "/blocks": ("GetBlocks", "resultGetBlocks", "[]BlockInfo"),
    "/block/{height}": ("GetBlockByHeight", "resultGetBlock", "*BlockInfo"),
    #"/block/{height}/validators": ("GetBlockValidators", "resultGetBlockValidators", "[]BlockValidator"),
    "/block/{height}/txs": ("GetBlockTransactions", "resultGetBlockTransactions", "[]TxInfo"),
    ### evm
    "/evm-contracts": ("GetEvmContracts", "resultGetEvmContracts", "[]EvmContract"),
    "/evm-transactions": ("GetEvmTransactions", "resultGetEvmTransactions", "[]EvmTransaction"),
    "/evm-accounts": ("GetEvmAccounts", "resultGetEvmAccounts", "[]EvmAccount"),
    "/evm-contracts/{address}": ("GetEvmContract", "resultGetEvmContract", "*EvmContract"),
    "/evm-transactions/{hash}": ("GetEvmTransaction", "resultGetEvmTransaction", "*EvmTransaction"),
    "/evm-accounts/{address}": ("GetEvmAccount", "resultGetEvmAccount", "*EvmAccount"),
    "/evm-contracts/{address}/transactions": ("GetEvmContractTransactions", "resultGetEvmContractTransactions", "[]EvmTransaction"),
    "/evm-contracts/{address}/events": ("GetEvmContractEvents", "resultGetEvmContractEvents", "[]EvmEvent"),
    "/evm-accounts/{address}/balances": ("GetEvmAccountBalances", "resultGetEvmAccountBalances", "[]EvmAccountBalance"),
    ### validators
    # kind = validator | candidate
    "/validators/{kind}": ("GetValidatorsByKind", "resultGetValidatorsByKind", "[]Validator"),
    #"/validators/{address}/coins": ("GetValidatorsCoins", "resultGetValidatorsCoins", "[]ValidatorStakedCoin"),
    "/validator/{address}": ("GetValidator", "resultGetValidator", "*Validator"),
    "/validator/{address}/stakes": ("GetValidatorStakes", "resultGetValidatorStakes", "[]ValidatorStake"),
    "/validator/{address}/stakes/nfts": ("GetValidatorStakesNFT", "resultGetValidatorStakesNFT", "[]ValidatorStakeNFT"),
}

def capitalize(s):
    if len(s)>1:
        return s[0].upper()+s[1:]
    elif len(s) == 1:
        return s[0].upper()
    else:
        return s

def type_conv(typename):
    if typename == "string":
        return "string"
    if typename == "number":
        return "uint64"
    if typename == "integer":
        return "int64"
    if typename == "boolean":
        return "bool"
    return typename

def formatting_directive(typename):
    if typename == "string":
        return "%s"
    if typename in ["uint64", "int64"]:
        return "%d"

def info2code(link, obj):
    """ info2code generates result type and function code
    We must declare output type and implement function converterXXX
    """
    methodName, responseType, resultType = ENDPOINTS_METHODS[link]
    comment = obj["summary"]
    if "description" in obj:
        comment = obj["description"]
    # params is for functions declaration
    params = []
    # paramsDescription is for function body
    paramsDescription = []
    optionalParams = []
    if "parameters" in obj:
        for p in obj["parameters"]:
            if "required" in p and p["required"]:
                params.append("%s %s" % (p["name"], type_conv(p["type"])))
                paramsDescription.append( (p["name"], type_conv(p["type"])))
            else:
                optionalParams.append(p["name"])
    useOptionalParams = False
    # now all optional params have limit and offset, but it may change
    if "limit" in optionalParams and "offset" in optionalParams:
        useOptionalParams = True
        params.append("opt *OptionalParams")

    outputType = responseType + " " + schema2type("", obj["responses"]["200"]["schema"], True)

    code = "///////////////\n\n" + \
        "type %s\n" % (outputType,) + \
        "// %s\n" % (link,) + \
        "// %s\n" % (comment,) + \
        "func (api *API) %s(%s) (%s, error) {\n" % (methodName, ", ".join(params), resultType) + \
        endpointMethodBody(methodName, link, paramsDescription, responseType, useOptionalParams) + \
        "}\n"

    verificationCode = "///////////////\n\n" + \
        "type %s\n" % (outputType,) + \
        "// %s\n" % (link,) + \
        "// %s\n" % (comment,) + \
        "func (api *API) Verification%s(%s) ([]string, error) {\n" % (methodName, ", ".join(params)) + \
        verificationMethodBody(methodName, link, paramsDescription, responseType, useOptionalParams) + \
        "}\n"

    readme_declaration = "%s(%s) (%s, error)" % (methodName, ", ".join(params), resultType)

    return code, verificationCode, readme_declaration

def schema2type(key, obj, skipJson=False):
    """ schema2type creates type declaration by swagger definition
    """
    if obj["type"] == "object":
        if "properties" in obj:
            return "%s struct {\n" % (capitalize(key),)+\
                "\n".join([schema2type(k,o) for k,o in obj["properties"].items()])+\
                ("\n}" if skipJson else "\n} `json:\"%s\"`" % (key,))
        else:
            return "%s interface{} `json:\"%s\"`" % (capitalize(key), key)
    if obj["type"] == "array":
        if obj["items"]["type"] == "object":
            return "%s []struct {\n" % (capitalize(key),)+\
                "\n".join([schema2type(k,o) for k,o in obj["items"]["properties"].items()])+\
                ("\n}" if skipJson else "\n} `json:\"%s\"`" % (key,))
        else:
            return "%s []%s `json:\"%s\"`" % (capitalize(key), type_conv(obj["items"]["type"]), key)
    return "%s %s `json:\"%s\"`" % (capitalize(key), type_conv(obj["type"]), key)

def endpointMethodBody(methodName, link, paramsDescription, responseType, useOptionalParams):
    newlink = 'var link = "'+link+'"\n'
    if len(paramsDescription)>0:
        newlink = "var r = strings.NewReplacer(\n"
        for name, typ in paramsDescription:
            newlink += '"{%s}", fmt.Sprintf("%s", %s),\n' % (name, formatting_directive(typ), name)
        newlink += ")\n"
        newlink += 'var link = r.Replace("'+link+'")\n'
    if useOptionalParams:
        newlink += "if opt != nil {\nlink += opt.String()\n}\n"
    return '''
    %s
    // request
    res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := %s{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})    
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
    // process result
    return converter%s(respValue)
    ''' % (newlink, responseType, methodName)

def verificationMethodBody(methodName, link, paramsDescription, responseType, useOptionalParams):
    newlink = 'var link = "'+link+'"\n'
    if len(paramsDescription)>0:
        newlink = "var r = strings.NewReplacer(\n"
        for name, typ in paramsDescription:
            newlink += '"{%s}", fmt.Sprintf("%s", %s),\n' % (name, formatting_directive(typ), name)
        newlink += ")\n"
        newlink += 'var link = r.Replace("'+link+'")\n'
    if useOptionalParams:
        newlink += "if opt != nil {\nlink += opt.String()\n}\n"
    return '''
    %s
    // request
    res, err := api.client.R().Get(link)
	if err != nil {
		return nil, err
	}
    if res.StatusCode() != 200 {
        return nil, fmt.Errorf("status: %%d", res.StatusCode())
    }

    // json decode
	respValue := %s{}
    err = json.Unmarshal(res.Body(), &respValue)
	if err != nil {
		return nil, err
	}
    // process result
    bz, _ := json.Marshal(respValue)
	if err != nil {
		return nil, err
	}

    result := []string{fmt.Sprintf("url: %%s", link)}
    result = append(result, compareJSON(bz, res.Body())...)
    return result, nil
    ''' % (newlink, responseType)

doc = json.load(open("documentation.json","r"))
readme = open("readme_appendix.md", "w")
endpointsFile = open("endpoints.go", "w")
verificationFile = open("verify_endpoints.go", "w")

endpointsFile.write("package api\n\n")
endpointsFile.write("// This file is autogenerated. DO NOT EDIT\n\n")
endpointsFile.write("import (\n\"fmt\"\n\"strings\"\n)\n")

verificationFile.write("package swagger\n\n")
verificationFile.write("// This file is autogenerated. DO NOT EDIT\n\n")
verificationFile.write("import (\n\"fmt\"\n\"strings\"\n\"encoding/json\"\n)\n")
for k in ENDPOINTS_METHODS.keys():
    print(k)
    if k not in doc["paths"]:
        print("!!! ",k," NOT FOUND")
        continue
    endpoint = doc["paths"][k]
    if "get" not in endpoint: continue
    code, verifyCode, readme_declaration = info2code(k, endpoint["get"])
    endpointsFile.write(code)
    endpointsFile.write("\n\n")
    verificationFile.write(verifyCode)
    verificationFile.write("\n\n")
    readme.write("- %s\n" % (readme_declaration,))
endpointsFile.close()
verificationFile.close()
readme.close()

print("Can be added:")
for k,endpoint in doc["paths"].items():
    if k in ENDPOINTS_METHODS: continue
    print(k)