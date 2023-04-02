package metricsengine

func (m *Metric) GetTVL(contractAddr string) (*int64, error) {
	tvl, err := m.transactionstorage.GetTVL(contractAddr)
	if err != nil {
		return nil, err
	}

	return tvl, nil
}
