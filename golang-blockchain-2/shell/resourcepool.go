package shell

type CpuResource struct {
	ResourceId     string
	ip             string
	Name           string
	Features       string
	Frequency      float64
	Cores          int
	ThreadsPerCore int
	LogicalCores   int
	Memory         float64
}

type GpuResource struct {
	ResourceId string
	ip         string
	Name       string
	Features   string
	Frequency  float64
	Memory     float64
}

type StorageResource struct {
	ResourceId string
	ip         string
	Port       int
	TotalSpace int64
	FreeSpace  int64
	Memory     int64
	Cores      int //no of cpus
}

// Fak stores private and public keys
type ResourcePool struct {
	CpuPool     []CpuResource
	GpuPool     []GpuResource
	StoragePool []StorageResource
	AdminPool   []CpuResource
}

func GetResourcePool() ResourcePool {
	tmp := ResourcePool{
		CpuPool: []CpuResource{
			CpuResource{
				ResourceId:     "res-cpu-0",
				ip:             "10.0.0.1",
				Name:           "AMD Ryzen 7",
				Features:       "VMX,SLA",
				Frequency:      3333.33,
				Cores:          4,
				ThreadsPerCore: 2,
				LogicalCores:   8,
				Memory:         4000000,
			},
			CpuResource{
				ResourceId:     "res-cpu-1",
				ip:             "20.0.0.2",
				Name:           "AMD Ryzen 5",
				Features:       "VMX,SLA",
				Frequency:      2666.3,
				Cores:          2,
				ThreadsPerCore: 2,
				LogicalCores:   4,
				Memory:         5000000,
			},
		},
		GpuPool: []GpuResource{
			GpuResource{
				ResourceId: "res-gpu-3",
				ip:         "33.0.0.33",
				Name:       "Zotac GTX 2066",
				Features:   "CUDA,BLAH,BLEH",
				Frequency:  2066.0,
				Memory:     2000000,
			},
			GpuResource{
				ResourceId: "res-gpu-4",
				ip:         "44.0.0.44",
				Name:       "NVidia GTX 3000Ti",
				Features:   "CUDA,BLAH,BLEH",
				Frequency:  3066.0,
				Memory:     3000000,
			},
		},
		StoragePool: []StorageResource{
			StorageResource{
				ResourceId: "res-storage-50",
				ip:         "55.0.0.5",
				Port:       5050,
				TotalSpace: 2000000000,
				FreeSpace:  1000000000,
				Memory:     2000000,
				Cores:      2,
			},
			StorageResource{
				ResourceId: "res-storage-51",
				ip:         "51.0.0.5",
				Port:       5050,
				TotalSpace: 2000000000,
				FreeSpace:  1000000000,
				Memory:     2000000,
				Cores:      2,
			},
		},
		AdminPool: []CpuResource{
			CpuResource{
				ResourceId:     "res-cpu-70",
				ip:             "70.0.0.1",
				Name:           "AMD Ryzen 7",
				Features:       "VMX,SLA",
				Frequency:      3666.66,
				Cores:          4,
				ThreadsPerCore: 2,
				LogicalCores:   8,
				Memory:         4000000,
			},
			CpuResource{
				ResourceId:     "res-cpu-71",
				ip:             "74.0.0.12",
				Name:           "Intel i5 7200u",
				Features:       "VMX,SLA",
				Frequency:      2666.3,
				Cores:          2,
				ThreadsPerCore: 2,
				LogicalCores:   4,
				Memory:         8000000,
			},
		},
	}
	return tmp
}
