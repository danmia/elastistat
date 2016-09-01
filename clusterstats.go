package main

type ClusterStats struct {
	ClusterName string `json:"cluster_name"`
	Indices     struct {
		Completion struct {
			Size        string `json:"size"`
			SizeInBytes int    `json:"size_in_bytes"`
		} `json:"completion"`
		Count int `json:"count"`
		Docs  struct {
			Count   int `json:"count"`
			Deleted int `json:"deleted"`
		} `json:"docs"`
		Fielddata struct {
			Evictions         int    `json:"evictions"`
			MemorySize        string `json:"memory_size"`
			MemorySizeInBytes int    `json:"memory_size_in_bytes"`
		} `json:"fielddata"`
		Percolate struct {
			Current           int    `json:"current"`
			MemorySize        string `json:"memory_size"`
			MemorySizeInBytes int    `json:"memory_size_in_bytes"`
			Queries           int    `json:"queries"`
			Time              string `json:"time"`
			TimeInMillis      int    `json:"time_in_millis"`
			Total             int    `json:"total"`
		} `json:"percolate"`
		QueryCache struct {
			CacheCount        int    `json:"cache_count"`
			CacheSize         int    `json:"cache_size"`
			Evictions         int    `json:"evictions"`
			HitCount          int    `json:"hit_count"`
			MemorySize        string `json:"memory_size"`
			MemorySizeInBytes int    `json:"memory_size_in_bytes"`
			MissCount         int    `json:"miss_count"`
			TotalCount        int    `json:"total_count"`
		} `json:"query_cache"`
		Segments struct {
			Count                       int    `json:"count"`
			DocValuesMemory             string `json:"doc_values_memory"`
			DocValuesMemoryInBytes      int    `json:"doc_values_memory_in_bytes"`
			FixedBitSet                 string `json:"fixed_bit_set"`
			FixedBitSetMemoryInBytes    int    `json:"fixed_bit_set_memory_in_bytes"`
			IndexWriterMaxMemory        string `json:"index_writer_max_memory"`
			IndexWriterMaxMemoryInBytes int    `json:"index_writer_max_memory_in_bytes"`
			IndexWriterMemory           string `json:"index_writer_memory"`
			IndexWriterMemoryInBytes    int    `json:"index_writer_memory_in_bytes"`
			Memory                      string `json:"memory"`
			MemoryInBytes               int    `json:"memory_in_bytes"`
			NormsMemory                 string `json:"norms_memory"`
			NormsMemoryInBytes          int    `json:"norms_memory_in_bytes"`
			StoredFieldsMemory          string `json:"stored_fields_memory"`
			StoredFieldsMemoryInBytes   int    `json:"stored_fields_memory_in_bytes"`
			TermVectorsMemory           string `json:"term_vectors_memory"`
			TermVectorsMemoryInBytes    int    `json:"term_vectors_memory_in_bytes"`
			TermsMemory                 string `json:"terms_memory"`
			TermsMemoryInBytes          int    `json:"terms_memory_in_bytes"`
			VersionMapMemory            string `json:"version_map_memory"`
			VersionMapMemoryInBytes     int    `json:"version_map_memory_in_bytes"`
		} `json:"segments"`
		Shards struct {
			Index struct {
				Primaries struct {
					Avg float64 `json:"avg"`
					Max int     `json:"max"`
					Min int     `json:"min"`
				} `json:"primaries"`
				Replication struct {
					Avg int `json:"avg"`
					Max int `json:"max"`
					Min int `json:"min"`
				} `json:"replication"`
				Shards struct {
					Avg float64 `json:"avg"`
					Max int     `json:"max"`
					Min int     `json:"min"`
				} `json:"shards"`
			} `json:"index"`
			Primaries   int     `json:"primaries"`
			Replication float64 `json:"replication"`
			Total       int     `json:"total"`
		} `json:"shards"`
		Store struct {
			Size                 string `json:"size"`
			SizeInBytes          int    `json:"size_in_bytes"`
			ThrottleTime         string `json:"throttle_time"`
			ThrottleTimeInMillis int    `json:"throttle_time_in_millis"`
		} `json:"store"`
	} `json:"indices"`
	Nodes struct {
		Count struct {
			Client     int `json:"client"`
			DataOnly   int `json:"data_only"`
			MasterData int `json:"master_data"`
			MasterOnly int `json:"master_only"`
			Total      int `json:"total"`
		} `json:"count"`
		Fs struct {
			Available        string `json:"available"`
			AvailableInBytes int    `json:"available_in_bytes"`
			Free             string `json:"free"`
			FreeInBytes      int    `json:"free_in_bytes"`
			Spins            string `json:"spins"`
			Total            string `json:"total"`
			TotalInBytes     int    `json:"total_in_bytes"`
		} `json:"fs"`
		Jvm struct {
			MaxUptime         string `json:"max_uptime"`
			MaxUptimeInMillis int    `json:"max_uptime_in_millis"`
			Mem               struct {
				HeapMax         string `json:"heap_max"`
				HeapMaxInBytes  int    `json:"heap_max_in_bytes"`
				HeapUsed        string `json:"heap_used"`
				HeapUsedInBytes int    `json:"heap_used_in_bytes"`
			} `json:"mem"`
			Threads  int `json:"threads"`
			Versions []struct {
				Count     int    `json:"count"`
				Version   string `json:"version"`
				VMName    string `json:"vm_name"`
				VMVendor  string `json:"vm_vendor"`
				VMVersion string `json:"vm_version"`
			} `json:"versions"`
		} `json:"jvm"`
		Os struct {
			AllocatedProcessors int `json:"allocated_processors"`
			AvailableProcessors int `json:"available_processors"`
			Mem                 struct {
				Total        string `json:"total"`
				TotalInBytes int    `json:"total_in_bytes"`
			} `json:"mem"`
			Names []struct {
				Count int    `json:"count"`
				Name  string `json:"name"`
			} `json:"names"`
		} `json:"os"`
		Plugins []struct {
			Description string `json:"description"`
			Jvm         bool   `json:"jvm"`
			Name        string `json:"name"`
			Site        bool   `json:"site"`
			URL         string `json:"url"`
			Version     string `json:"version"`
		} `json:"plugins"`
		Process struct {
			CPU struct {
				Percent int `json:"percent"`
			} `json:"cpu"`
			OpenFileDescriptors struct {
				Avg int `json:"avg"`
				Max int `json:"max"`
				Min int `json:"min"`
			} `json:"open_file_descriptors"`
		} `json:"process"`
		Versions []string `json:"versions"`
	} `json:"nodes"`
	Status    string `json:"status"`
	Timestamp int    `json:"timestamp"`
}

