package fundamentals

// Etf represents ETF fundamentals data.
type Etf struct {
	General struct {
		Code           string `json:"Code"`
		Type           string `json:"Type"`
		Name           string `json:"Name"`
		Exchange       string `json:"Exchange"`
		CurrencyCode   string `json:"CurrencyCode"`
		CurrencyName   string `json:"CurrencyName"`
		CurrencySymbol string `json:"CurrencySymbol"`
		CountryName    string `json:"CountryName"`
		CountryISO     string `json:"CountryISO"`
		OpenFigi       string `json:"OpenFigi"`
		Description    string `json:"Description"`
		Category       string `json:"Category"`
		UpdatedAt      string `json:"UpdatedAt"`
	} `json:"General"`
	Technicals struct {
		Beta          float64 `json:"Beta"`
		Five2WeekHigh float64 `json:"52WeekHigh"`
		Five2WeekLow  float64 `json:"52WeekLow"`
		Five0DayMA    float64 `json:"50DayMA"`
		Two00DayMA    float64 `json:"200DayMA"`
	} `json:"Technicals"`
	ETFData struct {
		Isin                    string `json:"ISIN"`
		CompanyName             string `json:"Company_Name"`
		CompanyURL              string `json:"Company_URL"`
		EtfURL                  string `json:"ETF_URL"`
		Domicile                string `json:"Domicile"`
		IndexName               string `json:"Index_Name"`
		Yield                   string `json:"Yield"`
		DividendPayingFrequency string `json:"Dividend_Paying_Frequency"`
		InceptionDate           string `json:"Inception_Date"`
		MaxAnnualMgmtCharge     string `json:"Max_Annual_Mgmt_Charge"`
		OngoingCharge           string `json:"Ongoing_Charge"`
		DateOngoingCharge       string `json:"Date_Ongoing_Charge"`
		NetExpenseRatio         string `json:"NetExpenseRatio"`
		AnnualHoldingsTurnover  string `json:"AnnualHoldingsTurnover"`
		TotalAssets             string `json:"TotalAssets"`
		AverageMktCapMil        string `json:"Average_Mkt_Cap_Mil"`
		MarketCapitalisation    struct {
			Mega   string `json:"Mega"`
			Big    string `json:"Big"`
			Medium string `json:"Medium"`
			Small  string `json:"Small"`
			Micro  string `json:"Micro"`
		} `json:"Market_Capitalisation"`
		AssetAllocation struct {
			Cash struct {
				Long      string `json:"Long_%"`
				Short     string `json:"Short_%"`
				NetAssets string `json:"Net_Assets_%"`
			} `json:"Cash"`
			NotClassified struct {
				Long      string `json:"Long_%"`
				Short     string `json:"Short_%"`
				NetAssets string `json:"Net_Assets_%"`
			} `json:"NotClassified"`
			StockNonUS struct {
				Long      string `json:"Long_%"`
				Short     string `json:"Short_%"`
				NetAssets string `json:"Net_Assets_%"`
			} `json:"Stock non-US"`
			Other struct {
				Long      string `json:"Long_%"`
				Short     string `json:"Short_%"`
				NetAssets string `json:"Net_Assets_%"`
			} `json:"Other"`
			StockUS struct {
				Long      string `json:"Long_%"`
				Short     string `json:"Short_%"`
				NetAssets string `json:"Net_Assets_%"`
			} `json:"Stock US"`
			Bond struct {
				Long      string `json:"Long_%"`
				Short     string `json:"Short_%"`
				NetAssets string `json:"Net_Assets_%"`
			} `json:"Bond"`
		} `json:"Asset_Allocation"`
		WorldRegions struct {
			NorthAmerica struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"North America"`
			UnitedKingdom struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"United Kingdom"`
			EuropeDeveloped struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Europe Developed"`
			EuropeEmerging struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Europe Emerging"`
			AfricaMiddleEast struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Africa/Middle East"`
			Japan struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Japan"`
			Australasia struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Australasia"`
			AsiaDeveloped struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Asia Developed"`
			AsiaEmerging struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Asia Emerging"`
			LatinAmerica struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Latin America"`
		} `json:"World_Regions"`
		SectorWeights struct {
			BasicMaterials struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Basic Materials"`
			ConsumerCyclicals struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Consumer Cyclicals"`
			FinancialServices struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Financial Services"`
			RealEstate struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Real Estate"`
			CommunicationServices struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Communication Services"`
			Energy struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Energy"`
			Industrials struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Industrials"`
			Technology struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Technology"`
			ConsumerDefensive struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Consumer Defensive"`
			Healthcare struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Healthcare"`
			Utilities struct {
				Equity             string `json:"Equity_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Utilities"`
		} `json:"Sector_Weights"`
		FixedIncome struct {
			EffectiveDuration struct {
				Fund               string `json:"Fund_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"EffectiveDuration"`
			ModifiedDuration struct {
				Fund               string `json:"Fund_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"ModifiedDuration"`
			EffectiveMaturity struct {
				Fund               string `json:"Fund_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"EffectiveMaturity"`
			CreditQuality struct {
				Fund               string `json:"Fund_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"CreditQuality"`
			Coupon struct {
				Fund               string `json:"Fund_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Coupon"`
			Price struct {
				Fund               string `json:"Fund_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"Price"`
			YieldToMaturity struct {
				Fund               string `json:"Fund_%"`
				RelativeToCategory string `json:"Relative_to_Category"`
			} `json:"YieldToMaturity"`
		} `json:"Fixed_Income"`
		HoldingsCount int `json:"Holdings_Count"`
		Top10Holdings map[string]struct {
			Code     string  `json:"Code"`
			Exchange string  `json:"Exchange"`
			Name     string  `json:"Name"`
			Sector   string  `json:"Sector"`
			Industry string  `json:"Industry"`
			Country  string  `json:"Country"`
			Region   string  `json:"Region"`
			Assets   float64 `json:"Assets_%"`
		} `json:"Top_10_Holdings"`
		Holdings map[string]struct {
			Code     string  `json:"Code"`
			Exchange string  `json:"Exchange"`
			Name     string  `json:"Name"`
			Sector   string  `json:"Sector"`
			Industry string  `json:"Industry"`
			Country  string  `json:"Country"`
			Region   string  `json:"Region"`
			Assets   float64 `json:"Assets_%"`
		} `json:"Holdings"`
		ValuationsGrowth struct {
			ValuationsRatesPortfolio struct {
				PriceProspectiveEarnings string `json:"Price/Prospective Earnings"`
				PriceBook                string `json:"Price/Book"`
				PriceSales               string `json:"Price/Sales"`
				PriceCashFlow            string `json:"Price/Cash Flow"`
				DividendYieldFactor      string `json:"Dividend-Yield Factor"`
			} `json:"Valuations_Rates_Portfolio"`
			ValuationsRatesToCategory struct {
				PriceProspectiveEarnings string `json:"Price/Prospective Earnings"`
				PriceBook                string `json:"Price/Book"`
				PriceSales               string `json:"Price/Sales"`
				PriceCashFlow            string `json:"Price/Cash Flow"`
				DividendYieldFactor      string `json:"Dividend-Yield Factor"`
			} `json:"Valuations_Rates_To_Category"`
			GrowthRatesPortfolio struct {
				LongTermProjectedEarningsGrowth string `json:"Long-Term Projected Earnings Growth"`
				HistoricalEarningsGrowth        string `json:"Historical Earnings Growth"`
				SalesGrowth                     string `json:"Sales Growth"`
				CashFlowGrowth                  string `json:"Cash-Flow Growth"`
				BookValueGrowth                 string `json:"Book-Value Growth"`
			} `json:"Growth_Rates_Portfolio"`
			GrowthRatesToCategory struct {
				LongTermProjectedEarningsGrowth string `json:"Long-Term Projected Earnings Growth"`
				HistoricalEarningsGrowth        string `json:"Historical Earnings Growth"`
				SalesGrowth                     string `json:"Sales Growth"`
				CashFlowGrowth                  string `json:"Cash-Flow Growth"`
				BookValueGrowth                 string `json:"Book-Value Growth"`
			} `json:"Growth_Rates_To_Category"`
		} `json:"Valuations_Growth"`
		MorningStar struct {
			Ratio               string `json:"Ratio"`
			CategoryBenchmark   string `json:"Category_Benchmark"`
			SustainabilityRatio string `json:"Sustainability_Ratio"`
		} `json:"MorningStar"`
		Performance struct {
			OneYVolatility   string `json:"1y_Volatility"`
			ThreeYVolatility string `json:"3y_Volatility"`
			ThreeYExpReturn  string `json:"3y_ExpReturn"`
			ThreeYSharpRatio string `json:"3y_SharpRatio"`
			ReturnsYTD       string `json:"Returns_YTD"`
			Returns1Y        string `json:"Returns_1Y"`
			Returns3Y        string `json:"Returns_3Y"`
			Returns5Y        string `json:"Returns_5Y"`
			Returns10Y       string `json:"Returns_10Y"`
		} `json:"Performance"`
	} `json:"ETF_Data"`
}
