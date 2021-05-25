package app

import (
	"encoding/json"
	"fmt"
	"ginShoppingMall/util"
	"github.com/gin-gonic/gin"
)

// GetProduct
// @Router /api/product [get]
// @Summary 商品信息
// @Description 商品信息
// @Tags API PRODUCT 商品
func GetProduct(c *gin.Context) {
	var mapResult map[string]interface{}
	jsonStr := `{
    "name": "喷绘-户内写真",
    "type": 2,
    "bn": "ABC0123",
    "store": 10000,
    "freez": 0,
    "description": "我是描述喷绘-户内写真",
    "spec": [
        {
            "type": "材质",
            "name": "户内PP纸无背胶",
            "child": [
                {
                    "type": "工艺",
                    "name": "进口高清户内写真",
                    "child": [
                        {
                            "type": "工艺&配件",
                            "name": "覆膜",
                            "part": [
                                {
                                    "type": "配件",
                                    "name": "亮膜"
                                },
                                {
                                    "type": "配件",
                                    "name": "哑膜"
                                },
                                {
                                    "type": "配件",
                                    "name": "地板膜（斜纹）"
                                },
                                {
                                    "type": "配件",
                                    "name": "地板膜（细纹）"
                                },
                                {
                                    "type": "配件",
                                    "name": "地板膜（粗纹）"
                                },
                                {
                                    "type": "配件",
                                    "name": "水晶膜"
                                },
                                {
                                    "type": "配件",
                                    "name": "双面胶（覆正面）"
                                },
                                {
                                    "type": "配件",
                                    "name": "双面胶（覆反面）"
                                },
                                {
                                    "type": "配件",
                                    "name": "可擦写膜"
                                },
                                {
                                    "type": "配件",
                                    "name": "黄底亮膜"
                                },
                                {
                                    "type": "配件",
                                    "name": "黄底哑膜"
                                }
                            ],
                            "child": [
                                {
                                    "type": "工艺",
                                    "name": "小块裁切"
                                },
                                {
                                    "type": "工艺",
                                    "name": "四角打扣",
                                    "child": [
                                        {
                                            "type": "工艺&配件",
                                            "name": "陪展架",
                                            "part": [
                                                {
                                                    "type": "配件",
                                                    "name": "X展架60普通型"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "X展架60加强型"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "X展架80普通型"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "X展架80加强型"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "门型展架60普通"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "门型展架80普通"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "门型展架80加强"
                                                }
                                            ]
                                        }
                                    ]
                                },
                                {
                                    "type": "工艺",
                                    "name": "异形切割"
                                },
                                {
                                    "type": "工艺&配件",
                                    "name": "陪挂轴",
                                    "part": [
                                        {
                                            "type": "配件",
                                            "name": "塑料挂轴"
                                        },
                                        {
                                            "type": "配件",
                                            "name": "铝合金挂轴"
                                        }
                                    ]
                                },
                                {
                                    "type": "工艺",
                                    "name": "易拉宝画面",
                                    "child": [
                                        {
                                            "type": "工艺&配件",
                                            "name": "易拉宝画面",
                                            "part": [
                                                {
                                                    "type": "配件",
                                                    "name": "易拉宝80塑钢"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "易拉宝80铝合金"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "易拉宝80豪华"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "易拉宝120铝合金"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "易拉宝120豪华"
                                                }
                                            ]
                                        }
                                    ]
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "type": "材质",
            "name": "户内PP背胶（户内写真）",
            "child": [
                {
                    "type": "工艺",
                    "name": "进口高清户内写真",
                    "child": [
                        {
                            "type": "工艺&配件",
                            "name": "覆膜",
                            "part": [
                                {
                                    "type": "配件",
                                    "name": "亮膜"
                                },
                                {
                                    "type": "配件",
                                    "name": "哑膜"
                                },
                                {
                                    "type": "配件",
                                    "name": "地板膜（斜纹）"
                                },
                                {
                                    "type": "配件",
                                    "name": "地板膜（细纹）"
                                },
                                {
                                    "type": "配件",
                                    "name": "地板膜（粗纹）"
                                },
                                {
                                    "type": "配件",
                                    "name": "水晶膜"
                                },
                                {
                                    "type": "配件",
                                    "name": "双面胶（覆正面）"
                                },
                                {
                                    "type": "配件",
                                    "name": "双面胶（覆反面）"
                                },
                                {
                                    "type": "配件",
                                    "name": "可擦写膜"
                                },
                                {
                                    "type": "配件",
                                    "name": "黄底亮膜"
                                },
                                {
                                    "type": "配件",
                                    "name": "黄底哑膜"
                                }
                            ],
                            "child": [
                                {
                                    "type": "工艺",
                                    "name": "小块裁切"
                                },
                                {
                                    "type": "工艺",
                                    "name": "双面对裱",
                                    "child": {
                                        "type": "工艺&配件",
                                        "name": "另面覆膜",
                                        "part": [
                                            {
                                                "type": "配件",
                                                "name": "亮膜"
                                            },
                                            {
                                                "type": "配件",
                                                "name": "哑膜"
                                            },
                                            {
                                                "type": "配件",
                                                "name": "地板膜（斜纹）"
                                            },
                                            {
                                                "type": "配件",
                                                "name": "地板膜（细纹）"
                                            },
                                            {
                                                "type": "配件",
                                                "name": "地板膜（粗纹）"
                                            },
                                            {
                                                "type": "配件",
                                                "name": "水晶膜"
                                            },
                                            {
                                                "type": "配件",
                                                "name": "双面胶（覆正面）"
                                            },
                                            {
                                                "type": "配件",
                                                "name": "双面胶（覆反面）"
                                            },
                                            {
                                                "type": "配件",
                                                "name": "可擦写膜"
                                            },
                                            {
                                                "type": "配件",
                                                "name": "黄底亮膜"
                                            },
                                            {
                                                "type": "配件",
                                                "name": "黄底哑膜"
                                            }
                                        ],
                                        "child": [
                                            {
                                                "type": "工艺",
                                                "name": "反面相同画面",
                                                "child": [
                                                    {
                                                        "type": "工艺",
                                                        "name": "包边",
                                                        "part": [
                                                            {
                                                                "type": "配件",
                                                                "name": "小灰边"
                                                            },
                                                            {
                                                                "type": "配件",
                                                                "name": "小银边"
                                                            },
                                                            {
                                                                "type": "配件",
                                                                "name": "中银边"
                                                            },
                                                            {
                                                                "type": "配件",
                                                                "name": "小黑边"
                                                            },
                                                            {
                                                                "type": "配件",
                                                                "name": "小金边"
                                                            },
                                                            {
                                                                "type": "配件",
                                                                "name": "中金边"
                                                            },
                                                            {
                                                                "type": "配件",
                                                                "name": "小蓝边"
                                                            }
                                                        ]
                                                    },
                                                    {
                                                        "type": "工艺",
                                                        "name": "四角打扣"
                                                    },
                                                    {
                                                        "type": "工艺",
                                                        "name": "正上方打单扣"
                                                    },
                                                    {
                                                        "type": "工艺",
                                                        "name": "正上方打双扣"
                                                    },
                                                    {
                                                        "type": "工艺",
                                                        "name": "异形切割"
                                                    },
                                                    {
                                                        "type": "工艺",
                                                        "name": "配挂轴",
                                                        "part": [
                                                            {
                                                                "type": "配件",
                                                                "name": "塑料挂轴"
                                                            },
                                                            {
                                                                "type": "配件",
                                                                "name": "铝合金挂轴"
                                                            }
                                                        ]
                                                    },
                                                    {
                                                        "type": "工艺",
                                                        "name": "配手举杆",
                                                        "part": [
                                                            {
                                                                "type": "配件",
                                                                "name": "手举杆塑料"
                                                            },
                                                            {
                                                                "type": "配件",
                                                                "name": "手举杆不锈钢"
                                                            }
                                                        ]
                                                    }
                                                ]
                                            },
                                            {
                                                "type": "工艺",
                                                "name": "反面不同画面"
                                            }
                                        ]
                                    }
                                },
                                {
                                    "type": "工艺",
                                    "name": "四角打扣",
                                    "child": [
                                        {
                                            "type": "工艺&配件",
                                            "name": "陪展架",
                                            "part": [
                                                {
                                                    "type": "配件",
                                                    "name": "X展架60普通型"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "X展架60加强型"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "X展架80普通型"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "X展架80加强型"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "门型展架60普通"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "门型展架80普通"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "门型展架80加强"
                                                }
                                            ]
                                        }
                                    ]
                                },
                                {
                                    "type": "工艺",
                                    "name": "异形切割"
                                },
                                {
                                    "type": "工艺&配件",
                                    "name": "陪挂轴",
                                    "part": [
                                        {
                                            "type": "配件",
                                            "name": "塑料挂轴"
                                        },
                                        {
                                            "type": "配件",
                                            "name": "铝合金挂轴"
                                        }
                                    ]
                                },
                                {
                                    "type": "工艺",
                                    "name": "易拉宝画面",
                                    "child": [
                                        {
                                            "type": "工艺&配件",
                                            "name": "易拉宝画面",
                                            "part": [
                                                {
                                                    "type": "配件",
                                                    "name": "易拉宝80塑钢"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "易拉宝80铝合金"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "易拉宝80豪华"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "易拉宝120铝合金"
                                                },
                                                {
                                                    "type": "配件",
                                                    "name": "易拉宝120豪华"
                                                }
                                            ]
                                        }
                                    ]
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "type": "材质",
            "name": "户内高光相纸"
        },
        {
            "type": "材质",
            "name": "户内饭喷灯片"
        },
        {
            "type": "材质",
            "name": "户内PVC硬片"
        },
        {
            "type": "材质",
            "name": "户内写真布"
        },
        {
            "type": "材质",
            "name": "水性白胶车贴"
        }
    ]
}`
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println(err.Error())
	}

	util.ResponseSuccess(c, mapResult)
}
