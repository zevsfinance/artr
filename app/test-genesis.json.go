// +build testing

package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var DefaultGenesisUsers map[string]sdk.AccAddress
var NonExistingUser sdk.AccAddress

func initDefaultGenesisUsers() {
	DefaultGenesisUsers = map[string]sdk.AccAddress{
		//                 1
		//         ┌───────┴───────┐
		//         2               3
		//     ┌───┴───┐       ┌───┴───┐
		//     4       5       6       7
		//  ┌──┴──┐ ┌──┴──┐ ┌──┴──┐ ┌──┴──┐
		//  8     9 A     B C     D E     F
		"user1":  accAddr("artr1d4ezqdj03uachct8hum0z9zlfftzdq2f6yzvhj"),
		"user2":  accAddr("artr1h8s8yf433ypjc5htavsyc9zvg3vk43vms03z3l"),
		"user3":  accAddr("artr1cjqvu8pns5ff3vcy4r7qwy57f2ts8chsjg8kyu"),
		"user4":  accAddr("artr1hdayszxl2ahw4rm0mct72rxzukq058mauc0dt6"),
		"user5":  accAddr("artr15n7wt45x4tkgunp25wylrjymjnkqug80a78jnn"),
		"user6":  accAddr("artr1sqh7ly9z3yme0k32f42qu330a663zkculmx7qq"),
		"user7":  accAddr("artr1uaz24ndash8umld4xpfn3zpknk5ske3xv4fqam"),
		"user8":  accAddr("artr14qn52aqd5dp4eycngm9e5ryrqdaf0u2zacp939"),
		"user9":  accAddr("artr1h2emj28qqj0e4k3azyzqdqznxdkf9r55w8qw60"),
		"user10": accAddr("artr1fedl94g9gqnntzqtgmxyp6msztzvw435nv2fr7"),
		"user11": accAddr("artr1j9j50h6k3v2p70nar6234etd2amxdeys8lzg9g"),
		"user12": accAddr("artr1tkl5vyca6mlfhl0zmjkl8nkcmlulpre58lpnwy"),
		"user13": accAddr("artr1xkwt5k2pktltp0jzk6hjz9k2k89l30448t74rq"),
		"user14": accAddr("artr1q2w5ytm97g490lcux69n3vfprqsdv65vtp8nun"),
		"user15": accAddr("artr1j29a9493fmlkjr9hmp54ltjun2meph9l5fhagf"),
		"root":   accAddr("artr1yhy6d3m4utltdml7w7zte7mqx5wyuskq9rr5vg"),
	}
	NonExistingUser = accAddr("artr1t8h48rk0wyvuvdae5aysmlnfly6rpqe68unx5k")
}
func accAddr(s string) sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(s)
	if err != nil {
		panic(err)
	}
	return addr
}

const DefaultUser1ConsPubKey = "artrvalconspub1zcjduepqpme87trszw7awc62ra2de9edwr40v7xy7yfhvpvds96fncagm04qxu308e"

//TODO: Cleanup
const defaultGenesis = `{
  "genesis_time": "2020-09-25T12:26:05.224653001Z",
  "chain_id": "test-cloud",
  "consensus_params": {
    "block": {
      "max_bytes": "22020096",
      "max_gas": "-1",
      "time_iota_ms": "1000"
    },
    "evidence": {
      "max_age_num_blocks": "100000",
      "max_age_duration": "172800000000000"
    },
    "validator": {
      "pub_key_types": [
        "ed25519"
      ]
    }
  },
  "app_hash": "",
  "app_state": {
    "auth": {
      "params": {
        "max_memo_characters": "256",
        "tx_sig_limit": "7",
        "tx_size_cost_per_byte": "10",
        "sig_verify_cost_ed25519": "590",
        "sig_verify_cost_secp256k1": "1000"
      },
      "accounts": [
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1yhy6d3m4utltdml7w7zte7mqx5wyuskq9rr5vg",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr14eyw3l9pszt7efjwvy6venvnhnaenn4uy8s9rk",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1k20rvph0j2pr4g3jwpprdaw23rathkxc2w6ce8",
            "coins": [
              {
                "denom": "uartr",
                "amount": "0"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },{
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1h93uunesjjcn2n8j47pq43ty5m7kusu0k39m7r",
            "coins": [
              {
                "denom": "uartr",
                "amount": "0"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },{
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1n2gkwynafyt6jqqjptyjeyzs4un6mvexf5vypg",
            "coins": [
              {
                "denom": "uartr",
                "amount": "0"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },{
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1kqv0kjz9g74zhk4hm8r7ac8m9d8ynkhgz9lekl",
            "coins": [
              {
                "denom": "uartr",
                "amount": "0"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },{
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr135qfyv227zu3cclmyycnc4u7f7vm86t80dvsw3",
            "coins": [
              {
                "denom": "uartr",
                "amount": "0"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },{
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1haatt0sqj2k0et9t4m4tp2qeeh8vvuf8zvur76",
            "coins": [
              {
                "denom": "uartr",
                "amount": "0"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
		{
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1d4ezqdj03uachct8hum0z9zlfftzdq2f6yzvhj",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1h8s8yf433ypjc5htavsyc9zvg3vk43vms03z3l",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              },
              {
                "denom": "uartrd",
                "amount": "10000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1cjqvu8pns5ff3vcy4r7qwy57f2ts8chsjg8kyu",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              },
              {
                "denom": "uartrd",
                "amount": "10000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1hdayszxl2ahw4rm0mct72rxzukq058mauc0dt6",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr15n7wt45x4tkgunp25wylrjymjnkqug80a78jnn",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1sqh7ly9z3yme0k32f42qu330a663zkculmx7qq",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1uaz24ndash8umld4xpfn3zpknk5ske3xv4fqam",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr14qn52aqd5dp4eycngm9e5ryrqdaf0u2zacp939",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1h2emj28qqj0e4k3azyzqdqznxdkf9r55w8qw60",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1fedl94g9gqnntzqtgmxyp6msztzvw435nv2fr7",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1j9j50h6k3v2p70nar6234etd2amxdeys8lzg9g",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1tkl5vyca6mlfhl0zmjkl8nkcmlulpre58lpnwy",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1xkwt5k2pktltp0jzk6hjz9k2k89l30448t74rq",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
        {
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1q2w5ytm97g490lcux69n3vfprqsdv65vtp8nun",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        },
	{
          "type": "cosmos-sdk/Account",
          "value": {
            "address": "artr1j29a9493fmlkjr9hmp54ltjun2meph9l5fhagf",
            "coins": [
              {
                "denom": "uartr",
                "amount": "1000000000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        }
      ]
    },
    "profile": {
      "params": {
        "creators": [
		   "artr1d4ezqdj03uachct8hum0z9zlfftzdq2f6yzvhj"
		],
		"card_magic": "112233592347"
      },
      "profiles": [
		{
			"address": "artr1yhy6d3m4utltdml7w7zte7mqx5wyuskq9rr5vg",
			"profile": {
				"nickname": "root",
				"active_unitl": "9000"
			}
		},
        {
	  "address": "artr1d4ezqdj03uachct8hum0z9zlfftzdq2f6yzvhj",
	  "profile": {
	    "nickname": "user1",
	    "active_until": "9000",
	    "noding": true,
	    "validator": true
	  }
	},
        {
	  "address": "artr1h8s8yf433ypjc5htavsyc9zvg3vk43vms03z3l",
	  "profile": {
	    "nickname": "user2",
	    "active_until": "9000"
	  }
	},
        {
	  "address": "artr1cjqvu8pns5ff3vcy4r7qwy57f2ts8chsjg8kyu",
	  "profile": {
	    "nickname": "user3",
	    "active_until": "9000"
	  }
	},
        {
	  "address": "artr1hdayszxl2ahw4rm0mct72rxzukq058mauc0dt6",
	  "profile": {
	    "nickname": "user4",
	    "active_until": "9000"
	  }
	},
        {
	  "address": "artr15n7wt45x4tkgunp25wylrjymjnkqug80a78jnn",
	  "profile": {
	    "nickname": "user5",
	    "active_until": "9000"
	  }
	},
        {
	  "address": "artr1sqh7ly9z3yme0k32f42qu330a663zkculmx7qq",
	  "profile": {
	    "nickname": "user6",
	    "active_until": "9000"
	  }
	},
        {
	  "address": "artr1uaz24ndash8umld4xpfn3zpknk5ske3xv4fqam",
	  "profile": {
	    "nickname": "user7",
	    "active_until": "9000"
	  }
	},
        {
	  "address": "artr14qn52aqd5dp4eycngm9e5ryrqdaf0u2zacp939",
	  "profile": {
	    "nickname": "user8",
	    "active_until": "9000"
	  }
	},
        {
	  "address": "artr1h2emj28qqj0e4k3azyzqdqznxdkf9r55w8qw60",
	  "profile": {
	    "nickname": "user9",
	    "active_until": "9000"
	  }
	},
        {
	  "address": "artr1fedl94g9gqnntzqtgmxyp6msztzvw435nv2fr7",
	  "profile": {
	    "nickname": "user10",
	    "active_until": "9000"
	  }
	},
        {
	  "address": "artr1j9j50h6k3v2p70nar6234etd2amxdeys8lzg9g",
	  "profile": {
	    "nickname": "user11",
	    "active_until": "9000"
	  }
	},
        {
	  "address": "artr1tkl5vyca6mlfhl0zmjkl8nkcmlulpre58lpnwy",
	  "profile": {
	    "nickname": "user12",
	    "active_until": "9000"
	  }
	},
        {
	  "address": "artr1xkwt5k2pktltp0jzk6hjz9k2k89l30448t74rq",
	  "profile": {
	    "nickname": "user13",
	    "active_until": "9000"
	  }
	},
        {
	  "address": "artr1q2w5ytm97g490lcux69n3vfprqsdv65vtp8nun",
	  "profile": {
	    "nickname": "user14",
	    "active_until": "9000"
	  }
	},
        {
	  "address": "artr1j29a9493fmlkjr9hmp54ltjun2meph9l5fhagf",
	  "profile": {
	    "nickname": "user15",
	    "active_until": "9000"
	  }
	}
      ] 
    },
    "noding": {
      "params": {
        "max_validators": 100,
        "jail_after": 2,
        "unjail_after": "120",
        "min_status": 2
      },
      "active": [
        {
          "account": "artr1d4ezqdj03uachct8hum0z9zlfftzdq2f6yzvhj",
          "pubkey": "artrvalconspub1zcjduepqpme87trszw7awc62ra2de9edwr40v7xy7yfhvpvds96fncagm04qxu308e"
        }
      ]
    },
    "schedule": {},
    "distribution": {
      "params": {
        "community_tax": "0.020000000000000000",
        "base_proposer_reward": "0.010000000000000000",
        "bonus_proposer_reward": "0.040000000000000000",
        "withdraw_addr_enabled": true
      },
      "fee_pool": {
        "community_pool": []
      },
      "delegator_withdraw_infos": [],
      "previous_proposer": "",
      "outstanding_rewards": [],
      "validator_accumulated_commissions": [],
      "validator_historical_rewards": [],
      "validator_current_rewards": [],
      "delegator_starting_infos": [],
      "validator_slash_events": []
    },
    "subscription": {
      "params": {
        "token_course": 100000,
        "subscription_price": 1990,
        "vpn_gb_price": 10,
        "storage_gb_price": 10,
        "base_vpn_gb": 7,
        "base_storage_gb": 5,
		"course_change_signers": [
		  "artr1d4ezqdj03uachct8hum0z9zlfftzdq2f6yzvhj"
		]
      },
      "activity": [
        {
		  "address": "artr1d4ezqdj03uachct8hum0z9zlfftzdq2f6yzvhj",
		  "renew_block": "9000",
		  "info": {
			"active": true,
			"expire_at": "9000"
		  }
		},
        {
	  "address": "artr1h8s8yf433ypjc5htavsyc9zvg3vk43vms03z3l",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	},
        {
	  "address": "artr1cjqvu8pns5ff3vcy4r7qwy57f2ts8chsjg8kyu",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	},
        {
	  "address": "artr1hdayszxl2ahw4rm0mct72rxzukq058mauc0dt6",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	},
        {
	  "address": "artr15n7wt45x4tkgunp25wylrjymjnkqug80a78jnn",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	},
        {
	  "address": "artr1sqh7ly9z3yme0k32f42qu330a663zkculmx7qq",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	},
        {
	  "address": "artr1uaz24ndash8umld4xpfn3zpknk5ske3xv4fqam",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	},
        {
	  "address": "artr14qn52aqd5dp4eycngm9e5ryrqdaf0u2zacp939",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	},
        {
	  "address": "artr1h2emj28qqj0e4k3azyzqdqznxdkf9r55w8qw60",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	},
        {
	  "address": "artr1fedl94g9gqnntzqtgmxyp6msztzvw435nv2fr7",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	},
        {
	  "address": "artr1j9j50h6k3v2p70nar6234etd2amxdeys8lzg9g",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	},
        {
	  "address": "artr1tkl5vyca6mlfhl0zmjkl8nkcmlulpre58lpnwy",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	},
        {
	  "address": "artr1xkwt5k2pktltp0jzk6hjz9k2k89l30448t74rq",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	},
        {
	  "address": "artr1q2w5ytm97g490lcux69n3vfprqsdv65vtp8nun",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	},
        {
	  "address": "artr1j29a9493fmlkjr9hmp54ltjun2meph9l5fhagf",
	  "renew_block": "9000",
	  "info": {
	    "active": true,
	    "expire_at": "9000"
	  }
	}
      ]
    },
    "referral": {
      "params": {
        "company_accounts": {
          "top_referrer": "artr1k20rvph0j2pr4g3jwpprdaw23rathkxc2w6ce8",
          "for_subscription": "artr1h93uunesjjcn2n8j47pq43ty5m7kusu0k39m7r",
          "promo_bonuses": "artr1n2gkwynafyt6jqqjptyjeyzs4un6mvexf5vypg",
          "status_bonuses": "artr1kqv0kjz9g74zhk4hm8r7ac8m9d8ynkhgz9lekl",
          "leader_bonuses": "artr135qfyv227zu3cclmyycnc4u7f7vm86t80dvsw3",
          "for_delegating": "artr1haatt0sqj2k0et9t4m4tp2qeeh8vvuf8zvur76"
        },
		"delegating_award": {
		  "company": "5/1000",
		  "network": ["5%", "1%", "1%", "2%", "1%", "1%", "1%", "1%", "1%", "5/1000"]
		},
		"subscription_award": {
		  "company": "10%",
		  "network": ["15%", "10%", "7%", "7%", "7%", "7%", "7%", "5%", "2%", "2%"]
		},
		"transition_cost": "10000000"
      },
      "top_level_accounts": [
        "artr1yhy6d3m4utltdml7w7zte7mqx5wyuskq9rr5vg"
      ],
      "other_accounts": [
	{
	  "referrer": "artr1yhy6d3m4utltdml7w7zte7mqx5wyuskq9rr5vg",
	  "referrals": [
            "artr1d4ezqdj03uachct8hum0z9zlfftzdq2f6yzvhj"
	  ]
	},
	{
	  "referrer": "artr1d4ezqdj03uachct8hum0z9zlfftzdq2f6yzvhj",
	  "referrals": [
            "artr1h8s8yf433ypjc5htavsyc9zvg3vk43vms03z3l",
	    "artr1cjqvu8pns5ff3vcy4r7qwy57f2ts8chsjg8kyu"
	  ]
	},
	{
          "referrer": "artr1h8s8yf433ypjc5htavsyc9zvg3vk43vms03z3l",
	  "referrals": [
            "artr1hdayszxl2ahw4rm0mct72rxzukq058mauc0dt6",
	    "artr15n7wt45x4tkgunp25wylrjymjnkqug80a78jnn"
	  ]
	},
	{
          "referrer": "artr1cjqvu8pns5ff3vcy4r7qwy57f2ts8chsjg8kyu",
	  "referrals": [
	    "artr1sqh7ly9z3yme0k32f42qu330a663zkculmx7qq",
	    "artr1uaz24ndash8umld4xpfn3zpknk5ske3xv4fqam"
	  ]
	},
	{
	  "referrer": "artr1hdayszxl2ahw4rm0mct72rxzukq058mauc0dt6",
	  "referrals": [
	    "artr14qn52aqd5dp4eycngm9e5ryrqdaf0u2zacp939",
	    "artr1h2emj28qqj0e4k3azyzqdqznxdkf9r55w8qw60"
	  ]
	},
	{
	  "referrer": "artr15n7wt45x4tkgunp25wylrjymjnkqug80a78jnn",
	  "referrals": [
	    "artr1fedl94g9gqnntzqtgmxyp6msztzvw435nv2fr7",
	    "artr1j9j50h6k3v2p70nar6234etd2amxdeys8lzg9g"
	  ]
	},
	{
	  "referrer": "artr1sqh7ly9z3yme0k32f42qu330a663zkculmx7qq",
	  "referrals": [
	    "artr1tkl5vyca6mlfhl0zmjkl8nkcmlulpre58lpnwy",
	    "artr1xkwt5k2pktltp0jzk6hjz9k2k89l30448t74rq"
	  ]
	},
	{
	  "referrer": "artr1uaz24ndash8umld4xpfn3zpknk5ske3xv4fqam",
	  "referrals": [
	    "artr1q2w5ytm97g490lcux69n3vfprqsdv65vtp8nun",
	    "artr1j29a9493fmlkjr9hmp54ltjun2meph9l5fhagf"
	  ]
	}
      ]
    },
    "voting": {
      "government": [
        "artr1d4ezqdj03uachct8hum0z9zlfftzdq2f6yzvhj",
		"artr1h8s8yf433ypjc5htavsyc9zvg3vk43vms03z3l",
		"artr1cjqvu8pns5ff3vcy4r7qwy57f2ts8chsjg8kyu"
      ],
      "params": {
        "voting_period": 2160
      }
    },
    "params": null,
    "artrbank": {
      "send_enabled": true
    },
    "vpn": {
      "params": {
        "signers": [
		  "artr1d4ezqdj03uachct8hum0z9zlfftzdq2f6yzvhj",
		  "artr1h8s8yf433ypjc5htavsyc9zvg3vk43vms03z3l"
		]
      },
      "vpn_statuses": null
    },
    "storage": {},
    "supply": {
      "supply": []
    },
    "delegating": {
      "params": {
        "percentage": {
          "minimal": "21",
          "thousand_plus": "24",
          "ten_k_plus": "27",
          "hundred_k_plus": "30"
        },
        "min_delegate": "1000"
      }
    },
    "earning": {
      "params": {
        "signers": [
		  "artr1d4ezqdj03uachct8hum0z9zlfftzdq2f6yzvhj",
		  "artr1h8s8yf433ypjc5htavsyc9zvg3vk43vms03z3l"
		]
      }
    }
  }
}`
