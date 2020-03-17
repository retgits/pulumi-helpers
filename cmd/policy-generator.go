package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Get the policy templates from the Serverless Application Repository
	url := "https://raw.githubusercontent.com/awslabs/serverless-application-model/develop/samtranslator/policy_templates_data/policy_templates.json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var policies map[string]interface{}
	json.Unmarshal(body, &policies)

	errPolicies := make([]string, 0)

	f, err := os.Create("../policies.go")
	if err != nil {
		panic(err)
	}

	f.WriteString(fmt.Sprintf("package policytemplates\n\nimport \"strings\"\n\n"))

	for key, val := range policies["Templates"].(map[string]interface{}) {
		pt := val.(map[string]interface{})

		name := key
		description := pt["Description"].(string)

		d := pt["Definition"].(map[string]interface{})
		s := d["Statement"].([]interface{})
		effect, action, resource, param, err := getPolicyStatement(s)
		if err != nil {
			errPolicies = append(errPolicies, name)
			f.WriteString(fmt.Sprintf("// Add%s %s\nfunc(f *Factory) Add%s() string {return ``}\n\n", name, description, name))
			continue
		}
		action = strings.ReplaceAll(action, ",", "\",\"")
		returnString := "f.policies = append(f.policies, policy)"
		if len(param) > 1 {
			returnString = fmt.Sprintf("policy = strings.ReplaceAll(policy, \"${%s}\", %s)\n%s", param, param, returnString)
			param = fmt.Sprintf("%s string", param)
		}
		f.WriteString(fmt.Sprintf("// Add%s %s\nfunc(f *Factory) Add%s(%s) {\npolicy := `{ \"Action\": [\"%s\"], \"Effect\": \"%s\", \"Resource\": \"%s\" }`\n%s\n}\n\n", name, description, name, param, action, effect, resource, returnString))
	}

	if len(errPolicies) > 0 {
		fmt.Printf("\n\nThere are %d policies that encountered errors:\n", len(errPolicies))
		for idx := range errPolicies {
			fmt.Println(errPolicies[idx])
		}
	}
}

func getPolicyStatement(p []interface{}) (string, string, string, string, error) {
	var err error

	policy := p[0].(map[string]interface{})

	effect := policy["Effect"].(string)

	action := ""
	switch v := policy["Action"].(type) {
	case string:
		action = policy["Action"].(string)
	case []interface{}:
		a := policy["Action"].([]interface{})
		actions := make([]string, len(a))
		for idx := range actions {
			actions[idx] = a[idx].(string)
		}
		action = strings.Join(actions, ",")
	default:
		err = fmt.Errorf("unknown action type %T in getPolicyStatement", v)
	}

	resource := ""
	param := ""
	switch v := policy["Resource"].(type) {
	case string:
		resource = policy["Resource"].(string)
	case map[string]interface{}:
		res := policy["Resource"].(map[string]interface{})
		resource, param, err = getResource(res["Fn::Sub"])
	default:
		err = fmt.Errorf("unknown resource type %T in getPolicyStatement", v)
	}

	return effect, action, resource, param, err
}

func getResource(p interface{}) (string, string, error) {
	switch v := p.(type) {
	case string:
		return p.(string), "", nil
	case []interface{}:
		a := p.([]interface{})
		param := getParameter(a[1].(map[string]interface{}))
		return a[0].(string), param, nil
	default:
		return "", "", fmt.Errorf("unknown resource type %T in getResource", v)
	}
}

func getParameter(p map[string]interface{}) string {
	parameters := make([]string, 0)
	for key := range p {
		parameters = append(parameters, key)
	}

	var parameterString string

	switch len(parameters) {
	case 0:
		parameterString = ""
	case 1:
		parameterString = parameters[0]
	}

	return parameterString
}
