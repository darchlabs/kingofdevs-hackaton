package metricsengine

func (m *Metric) GetTotalAddresses(contractAddr string) (*int64, error) {
	total, err := m.transactionstorage.ListTotalAddresses(contractAddr)
	if err != nil {
		return nil, err
	}

	return total, nil

}
