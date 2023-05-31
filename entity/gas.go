package entity

type YungoGas struct {
	Gasfil	float64  `json:"gasfil" xorm:"comment('用户名') DOUBLE"`
	Pleagefil float64 `json:"pleagefil" xorm:"comment('用户名') DOUBLE"`
	Totalfil float64 `json:"totalfil" xorm:"comment('用户名') DOUBLE"`
	Cnytofil float64 `json:"totalfil" xorm:"comment('用户名') bigint"`
	CreateTime int64 `json:"totalfil" xorm:"comment('用户名') bigint"`
}

//func (gas *GasResult)GetGasFIL(pc2,c2 *api.InvocResult)(result float64){
//	//pc2gas费
//	f099 := *pc2.GasCost.TotalCost.Int
//	total := f099.Add(&f099,pc2.GasCost.MinerTip.Int)
//
//	//C2gas费
//	f := *c2.GasCost.TotalCost.Int
//	totalc2 := f099.Add(&f,c2.GasCost.MinerTip.Int)
//
//	count := total.Add(total,totalc2).String()
//	result = utils.NanoOrAttoToFIL(count,utils.AttoFIL)
//
//	return
//}
////质押费
//func (gas *GasResult)GetPleageFIL(pc2,c2 *api.InvocResult)(result float64){
//
//	pc2pledge := *pc2.Msg.Value.Int
//
//	count := pc2pledge.Add(&pc2pledge,c2.Msg.Value.Int).String()
//	result = utils.NanoOrAttoToFIL(count,utils.AttoFIL)
//	return
//}
//
//func (gas *GasResult)GetTotalFIL(pc2,c2 *api.InvocResult)(result float64){
//	//pc2gas费
//	f099 := *pc2.GasCost.TotalCost.Int
//	total := f099.Add(&f099,pc2.GasCost.MinerTip.Int)
//	fmt.Println("f099:",f099,",pc2.GasCost.TotalCost.Int:",pc2.GasCost.TotalCost.Int)
//	//C2gas费
//	f099 = *c2.GasCost.TotalCost.Int
//	totalc2 := f099.Add(&f099,c2.GasCost.MinerTip.Int)
//
//	gascount := total.Add(total,totalc2)
//
//	//质押费
//	pc2pledge := *pc2.Msg.Value.Int
//	pledgecount := pc2pledge.Add(&pc2pledge,c2.Msg.Value.Int)
//
//	//总和
//	count := gascount.Add(gascount,pledgecount).String()
//	result = utils.NanoOrAttoToFIL(count,utils.AttoFIL)
//
//
//	return
//}
//
//func (gas *GasResult)GetGasResul(pc2,c2 *api.InvocResult){
//	gas.GasFIL = gas.GetGasFIL(pc2,c2)
//	gas.PleageFIL = gas.GetPleageFIL(pc2,c2)
//	gas.TotalFIL = gas.GetTotalFIL(pc2,c2)
//}
