package shell

// Fak stores private and public keys
//type Pallet struct {
//	User_Id           []byte
//	Function_Key      []byte
//	Amount            float64
//	Timestamp         int64
//	RunTime           Runtime
//	RequiredResourses RequiredResourses
//}

type Palette struct {
	User_id       string
	Function_name string
	//	Function_key                      string
	//	Function_key_permissions          string
	//	Function_key_permissions_required string
	//	Time_stamp                        string
	//	Sender_address                    string
	//	Receiver_address                  string
	//	Amount                            float64
	//	Arbitrary_data                    string
	//	Validator_type                    string
	//	Nonce                             string
	//	Network_requirements              NetworkRequirement
	Storage_requirements StorageRequirement
	Cpu_requirements     CpuRequirement
	Gpu_Requirements     GpuRequirement
}

type Runtime struct {
	Language string
	code     string
}

type Shell struct {
	Hash                    string
	Palette                 Palette
	Orchestration_key       string
	Computational_resources ResourcePool
}

type RequiredResourses struct {
	Cpu     CpuRequirement
	Gpu     GpuRequirement
	Storage StorageRequirement
}

type CpuRequirement struct {
	Count        int
	Memory       int64
	Features     string
	MinFrequency int64
}

type GpuRequirement struct {
	Count        int
	Memory       int64
	Features     string
	MinFrequency int64
}

type StorageRequirement struct {
	Type      string
	Count     int
	Space     int
	Hydration string
}

type NetworkRequirement struct {
	Complexity     string
	Validator_type string
}
