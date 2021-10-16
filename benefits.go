package bamboohr

import "fmt"

// ListBenefitDeductionTypes lists the types of benefit deductions the company has
func (b *Client) ListBenefitDeductionTypes() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefits/settings/deduction_types/all", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// ListBenefitCoverage lists the benefit coverages
func (b *Client) ListBenefitCoverage() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitcoverages", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// GetBenefitCoverage gets the benefit coverage for a specific employee
func (b *Client) GetBenefitCoverage(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitcoverage/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// GetEmployeeDependent gets an employee dependent by ID
func (b *Client) GetEmployeeDependent(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/employeedependents/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// ListEmployeeDependents lists an employees dependents
func (b *Client) ListEmployeeDependents(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/employeedependents?employeeid=%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// ListBenefitPlans lists the company benefit plans
func (b *Client) ListBenefitPlans() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitplans", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// GetBenefitPlan gets a specific benefit plan
func (b *Client) GetBenefitPlan(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitplans/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// GetBenefitPlanCoverage gets the coverage of a benefit plan
func (b *Client) GetBenefitPlanCoverage(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitplancoverages/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// ListBenefitPlanCoverages lists the coverages for all benefit plans
func (b *Client) ListBenefitPlanCoverages() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitplancoverages", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// ListBenefitGroups lists the benefit groups
func (b *Client) ListBenefitGroups() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitgroups", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// GetBenefitGroup gets a specific benefit group
func (b *Client) GetBenefitGroup(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitgroups/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// ListBenefitGroupEmployees lists the employees for benefit groups
func (b *Client) ListBenefitGroupEmployees() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitgroupemployees", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// GetBenefitGroupEmployee gets a benefit group employee
func (b *Client) GetBenefitGroupEmployee(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitgroupemployees/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// ListBenefitGroupPlans lists the benefit group plans in your company
func (b *Client) ListBenefitGroupPlans() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitgroupplans", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// GetBenefitGroupPlan gets a specific benefit group plan
func (b *Client) GetBenefitGroupPlan(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitgroupplans/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// ListBenefitGroupPlanCosts lists the costs of benefit plans
func (b *Client) ListBenefitGroupPlanCosts() ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitgroupplancosts", b.APIEndpoint)

	return b.getRequest(endpointURL)
}

// GetBenefitGroupPlanCost gets the cost of a specific benefit plan
func (b *Client) GetBenefitGroupPlanCost(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/benefitgroupplancosts/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// GetEmployeeDeductionsByBenefitPlan gets the deductions by a plan
func (b *Client) GetEmployeeDeductionsByBenefitPlan(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/employee/deductions/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// GetBenefitPlanDeductionsByEmployee gets the deductions by an employee
func (b *Client) GetBenefitPlanDeductionsByEmployee(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/employee/plans/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}

// GetBenefitDeductionsForEmployee gets the benefit deductions for an employee
func (b *Client) GetBenefitDeductionsForEmployee(id string) ([]byte, error) {
	endpointURL := fmt.Sprintf("%s/payroll/deductions/%s", b.APIEndpoint, id)

	return b.getRequest(endpointURL)
}
