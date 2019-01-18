package gomcbot

var blockStates = `
{
  "minecraft:air": {
    "states": [
      {
        "id": 0,
        "default": true
      }
    ]
  },
  "minecraft:stone": {
    "states": [
      {
        "id": 1,
        "default": true
      }
    ]
  },
  "minecraft:granite": {
    "states": [
      {
        "id": 2,
        "default": true
      }
    ]
  },
  "minecraft:polished_granite": {
    "states": [
      {
        "id": 3,
        "default": true
      }
    ]
  },
  "minecraft:diorite": {
    "states": [
      {
        "id": 4,
        "default": true
      }
    ]
  },
  "minecraft:polished_diorite": {
    "states": [
      {
        "id": 5,
        "default": true
      }
    ]
  },
  "minecraft:andesite": {
    "states": [
      {
        "id": 6,
        "default": true
      }
    ]
  },
  "minecraft:polished_andesite": {
    "states": [
      {
        "id": 7,
        "default": true
      }
    ]
  },
  "minecraft:grass_block": {
    "properties": {
      "snowy": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "snowy": "true"
        },
        "id": 8
      },
      {
        "properties": {
          "snowy": "false"
        },
        "id": 9,
        "default": true
      }
    ]
  },
  "minecraft:dirt": {
    "states": [
      {
        "id": 10,
        "default": true
      }
    ]
  },
  "minecraft:coarse_dirt": {
    "states": [
      {
        "id": 11,
        "default": true
      }
    ]
  },
  "minecraft:podzol": {
    "properties": {
      "snowy": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "snowy": "true"
        },
        "id": 12
      },
      {
        "properties": {
          "snowy": "false"
        },
        "id": 13,
        "default": true
      }
    ]
  },
  "minecraft:cobblestone": {
    "states": [
      {
        "id": 14,
        "default": true
      }
    ]
  },
  "minecraft:oak_planks": {
    "states": [
      {
        "id": 15,
        "default": true
      }
    ]
  },
  "minecraft:spruce_planks": {
    "states": [
      {
        "id": 16,
        "default": true
      }
    ]
  },
  "minecraft:birch_planks": {
    "states": [
      {
        "id": 17,
        "default": true
      }
    ]
  },
  "minecraft:jungle_planks": {
    "states": [
      {
        "id": 18,
        "default": true
      }
    ]
  },
  "minecraft:acacia_planks": {
    "states": [
      {
        "id": 19,
        "default": true
      }
    ]
  },
  "minecraft:dark_oak_planks": {
    "states": [
      {
        "id": 20,
        "default": true
      }
    ]
  },
  "minecraft:oak_sapling": {
    "properties": {
      "stage": [
        "0",
        "1"
      ]
    },
    "states": [
      {
        "properties": {
          "stage": "0"
        },
        "id": 21,
        "default": true
      },
      {
        "properties": {
          "stage": "1"
        },
        "id": 22
      }
    ]
  },
  "minecraft:spruce_sapling": {
    "properties": {
      "stage": [
        "0",
        "1"
      ]
    },
    "states": [
      {
        "properties": {
          "stage": "0"
        },
        "id": 23,
        "default": true
      },
      {
        "properties": {
          "stage": "1"
        },
        "id": 24
      }
    ]
  },
  "minecraft:birch_sapling": {
    "properties": {
      "stage": [
        "0",
        "1"
      ]
    },
    "states": [
      {
        "properties": {
          "stage": "0"
        },
        "id": 25,
        "default": true
      },
      {
        "properties": {
          "stage": "1"
        },
        "id": 26
      }
    ]
  },
  "minecraft:jungle_sapling": {
    "properties": {
      "stage": [
        "0",
        "1"
      ]
    },
    "states": [
      {
        "properties": {
          "stage": "0"
        },
        "id": 27,
        "default": true
      },
      {
        "properties": {
          "stage": "1"
        },
        "id": 28
      }
    ]
  },
  "minecraft:acacia_sapling": {
    "properties": {
      "stage": [
        "0",
        "1"
      ]
    },
    "states": [
      {
        "properties": {
          "stage": "0"
        },
        "id": 29,
        "default": true
      },
      {
        "properties": {
          "stage": "1"
        },
        "id": 30
      }
    ]
  },
  "minecraft:dark_oak_sapling": {
    "properties": {
      "stage": [
        "0",
        "1"
      ]
    },
    "states": [
      {
        "properties": {
          "stage": "0"
        },
        "id": 31,
        "default": true
      },
      {
        "properties": {
          "stage": "1"
        },
        "id": 32
      }
    ]
  },
  "minecraft:bedrock": {
    "states": [
      {
        "id": 33,
        "default": true
      }
    ]
  },
  "minecraft:water": {
    "properties": {
      "level": [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "10",
        "11",
        "12",
        "13",
        "14",
        "15"
      ]
    },
    "states": [
      {
        "properties": {
          "level": "0"
        },
        "id": 34,
        "default": true
      },
      {
        "properties": {
          "level": "1"
        },
        "id": 35
      },
      {
        "properties": {
          "level": "2"
        },
        "id": 36
      },
      {
        "properties": {
          "level": "3"
        },
        "id": 37
      },
      {
        "properties": {
          "level": "4"
        },
        "id": 38
      },
      {
        "properties": {
          "level": "5"
        },
        "id": 39
      },
      {
        "properties": {
          "level": "6"
        },
        "id": 40
      },
      {
        "properties": {
          "level": "7"
        },
        "id": 41
      },
      {
        "properties": {
          "level": "8"
        },
        "id": 42
      },
      {
        "properties": {
          "level": "9"
        },
        "id": 43
      },
      {
        "properties": {
          "level": "10"
        },
        "id": 44
      },
      {
        "properties": {
          "level": "11"
        },
        "id": 45
      },
      {
        "properties": {
          "level": "12"
        },
        "id": 46
      },
      {
        "properties": {
          "level": "13"
        },
        "id": 47
      },
      {
        "properties": {
          "level": "14"
        },
        "id": 48
      },
      {
        "properties": {
          "level": "15"
        },
        "id": 49
      }
    ]
  },
  "minecraft:lava": {
    "properties": {
      "level": [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "10",
        "11",
        "12",
        "13",
        "14",
        "15"
      ]
    },
    "states": [
      {
        "properties": {
          "level": "0"
        },
        "id": 50,
        "default": true
      },
      {
        "properties": {
          "level": "1"
        },
        "id": 51
      },
      {
        "properties": {
          "level": "2"
        },
        "id": 52
      },
      {
        "properties": {
          "level": "3"
        },
        "id": 53
      },
      {
        "properties": {
          "level": "4"
        },
        "id": 54
      },
      {
        "properties": {
          "level": "5"
        },
        "id": 55
      },
      {
        "properties": {
          "level": "6"
        },
        "id": 56
      },
      {
        "properties": {
          "level": "7"
        },
        "id": 57
      },
      {
        "properties": {
          "level": "8"
        },
        "id": 58
      },
      {
        "properties": {
          "level": "9"
        },
        "id": 59
      },
      {
        "properties": {
          "level": "10"
        },
        "id": 60
      },
      {
        "properties": {
          "level": "11"
        },
        "id": 61
      },
      {
        "properties": {
          "level": "12"
        },
        "id": 62
      },
      {
        "properties": {
          "level": "13"
        },
        "id": 63
      },
      {
        "properties": {
          "level": "14"
        },
        "id": 64
      },
      {
        "properties": {
          "level": "15"
        },
        "id": 65
      }
    ]
  },
  "minecraft:sand": {
    "states": [
      {
        "id": 66,
        "default": true
      }
    ]
  },
  "minecraft:red_sand": {
    "states": [
      {
        "id": 67,
        "default": true
      }
    ]
  },
  "minecraft:gravel": {
    "states": [
      {
        "id": 68,
        "default": true
      }
    ]
  },
  "minecraft:gold_ore": {
    "states": [
      {
        "id": 69,
        "default": true
      }
    ]
  },
  "minecraft:iron_ore": {
    "states": [
      {
        "id": 70,
        "default": true
      }
    ]
  },
  "minecraft:coal_ore": {
    "states": [
      {
        "id": 71,
        "default": true
      }
    ]
  },
  "minecraft:oak_log": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 72
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 73,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 74
      }
    ]
  },
  "minecraft:spruce_log": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 75
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 76,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 77
      }
    ]
  },
  "minecraft:birch_log": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 78
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 79,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 80
      }
    ]
  },
  "minecraft:jungle_log": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 81
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 82,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 83
      }
    ]
  },
  "minecraft:acacia_log": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 84
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 85,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 86
      }
    ]
  },
  "minecraft:dark_oak_log": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 87
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 88,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 89
      }
    ]
  },
  "minecraft:stripped_spruce_log": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 90
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 91,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 92
      }
    ]
  },
  "minecraft:stripped_birch_log": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 93
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 94,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 95
      }
    ]
  },
  "minecraft:stripped_jungle_log": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 96
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 97,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 98
      }
    ]
  },
  "minecraft:stripped_acacia_log": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 99
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 100,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 101
      }
    ]
  },
  "minecraft:stripped_dark_oak_log": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 102
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 103,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 104
      }
    ]
  },
  "minecraft:stripped_oak_log": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 105
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 106,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 107
      }
    ]
  },
  "minecraft:oak_wood": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 108
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 109,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 110
      }
    ]
  },
  "minecraft:spruce_wood": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 111
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 112,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 113
      }
    ]
  },
  "minecraft:birch_wood": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 114
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 115,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 116
      }
    ]
  },
  "minecraft:jungle_wood": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 117
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 118,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 119
      }
    ]
  },
  "minecraft:acacia_wood": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 120
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 121,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 122
      }
    ]
  },
  "minecraft:dark_oak_wood": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 123
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 124,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 125
      }
    ]
  },
  "minecraft:stripped_oak_wood": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 126
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 127,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 128
      }
    ]
  },
  "minecraft:stripped_spruce_wood": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 129
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 130,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 131
      }
    ]
  },
  "minecraft:stripped_birch_wood": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 132
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 133,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 134
      }
    ]
  },
  "minecraft:stripped_jungle_wood": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 135
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 136,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 137
      }
    ]
  },
  "minecraft:stripped_acacia_wood": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 138
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 139,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 140
      }
    ]
  },
  "minecraft:stripped_dark_oak_wood": {
    "properties": {
      "axis": [
        "x",
        "y",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 141
      },
      {
        "properties": {
          "axis": "y"
        },
        "id": 142,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 143
      }
    ]
  },
  "minecraft:oak_leaves": {
    "properties": {
      "distance": [
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7"
      ],
      "persistent": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "distance": "1",
          "persistent": "true"
        },
        "id": 144
      },
      {
        "properties": {
          "distance": "1",
          "persistent": "false"
        },
        "id": 145
      },
      {
        "properties": {
          "distance": "2",
          "persistent": "true"
        },
        "id": 146
      },
      {
        "properties": {
          "distance": "2",
          "persistent": "false"
        },
        "id": 147
      },
      {
        "properties": {
          "distance": "3",
          "persistent": "true"
        },
        "id": 148
      },
      {
        "properties": {
          "distance": "3",
          "persistent": "false"
        },
        "id": 149
      },
      {
        "properties": {
          "distance": "4",
          "persistent": "true"
        },
        "id": 150
      },
      {
        "properties": {
          "distance": "4",
          "persistent": "false"
        },
        "id": 151
      },
      {
        "properties": {
          "distance": "5",
          "persistent": "true"
        },
        "id": 152
      },
      {
        "properties": {
          "distance": "5",
          "persistent": "false"
        },
        "id": 153
      },
      {
        "properties": {
          "distance": "6",
          "persistent": "true"
        },
        "id": 154
      },
      {
        "properties": {
          "distance": "6",
          "persistent": "false"
        },
        "id": 155
      },
      {
        "properties": {
          "distance": "7",
          "persistent": "true"
        },
        "id": 156
      },
      {
        "properties": {
          "distance": "7",
          "persistent": "false"
        },
        "id": 157,
        "default": true
      }
    ]
  },
  "minecraft:spruce_leaves": {
    "properties": {
      "distance": [
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7"
      ],
      "persistent": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "distance": "1",
          "persistent": "true"
        },
        "id": 158
      },
      {
        "properties": {
          "distance": "1",
          "persistent": "false"
        },
        "id": 159
      },
      {
        "properties": {
          "distance": "2",
          "persistent": "true"
        },
        "id": 160
      },
      {
        "properties": {
          "distance": "2",
          "persistent": "false"
        },
        "id": 161
      },
      {
        "properties": {
          "distance": "3",
          "persistent": "true"
        },
        "id": 162
      },
      {
        "properties": {
          "distance": "3",
          "persistent": "false"
        },
        "id": 163
      },
      {
        "properties": {
          "distance": "4",
          "persistent": "true"
        },
        "id": 164
      },
      {
        "properties": {
          "distance": "4",
          "persistent": "false"
        },
        "id": 165
      },
      {
        "properties": {
          "distance": "5",
          "persistent": "true"
        },
        "id": 166
      },
      {
        "properties": {
          "distance": "5",
          "persistent": "false"
        },
        "id": 167
      },
      {
        "properties": {
          "distance": "6",
          "persistent": "true"
        },
        "id": 168
      },
      {
        "properties": {
          "distance": "6",
          "persistent": "false"
        },
        "id": 169
      },
      {
        "properties": {
          "distance": "7",
          "persistent": "true"
        },
        "id": 170
      },
      {
        "properties": {
          "distance": "7",
          "persistent": "false"
        },
        "id": 171,
        "default": true
      }
    ]
  },
  "minecraft:birch_leaves": {
    "properties": {
      "distance": [
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7"
      ],
      "persistent": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "distance": "1",
          "persistent": "true"
        },
        "id": 172
      },
      {
        "properties": {
          "distance": "1",
          "persistent": "false"
        },
        "id": 173
      },
      {
        "properties": {
          "distance": "2",
          "persistent": "true"
        },
        "id": 174
      },
      {
        "properties": {
          "distance": "2",
          "persistent": "false"
        },
        "id": 175
      },
      {
        "properties": {
          "distance": "3",
          "persistent": "true"
        },
        "id": 176
      },
      {
        "properties": {
          "distance": "3",
          "persistent": "false"
        },
        "id": 177
      },
      {
        "properties": {
          "distance": "4",
          "persistent": "true"
        },
        "id": 178
      },
      {
        "properties": {
          "distance": "4",
          "persistent": "false"
        },
        "id": 179
      },
      {
        "properties": {
          "distance": "5",
          "persistent": "true"
        },
        "id": 180
      },
      {
        "properties": {
          "distance": "5",
          "persistent": "false"
        },
        "id": 181
      },
      {
        "properties": {
          "distance": "6",
          "persistent": "true"
        },
        "id": 182
      },
      {
        "properties": {
          "distance": "6",
          "persistent": "false"
        },
        "id": 183
      },
      {
        "properties": {
          "distance": "7",
          "persistent": "true"
        },
        "id": 184
      },
      {
        "properties": {
          "distance": "7",
          "persistent": "false"
        },
        "id": 185,
        "default": true
      }
    ]
  },
  "minecraft:jungle_leaves": {
    "properties": {
      "distance": [
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7"
      ],
      "persistent": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "distance": "1",
          "persistent": "true"
        },
        "id": 186
      },
      {
        "properties": {
          "distance": "1",
          "persistent": "false"
        },
        "id": 187
      },
      {
        "properties": {
          "distance": "2",
          "persistent": "true"
        },
        "id": 188
      },
      {
        "properties": {
          "distance": "2",
          "persistent": "false"
        },
        "id": 189
      },
      {
        "properties": {
          "distance": "3",
          "persistent": "true"
        },
        "id": 190
      },
      {
        "properties": {
          "distance": "3",
          "persistent": "false"
        },
        "id": 191
      },
      {
        "properties": {
          "distance": "4",
          "persistent": "true"
        },
        "id": 192
      },
      {
        "properties": {
          "distance": "4",
          "persistent": "false"
        },
        "id": 193
      },
      {
        "properties": {
          "distance": "5",
          "persistent": "true"
        },
        "id": 194
      },
      {
        "properties": {
          "distance": "5",
          "persistent": "false"
        },
        "id": 195
      },
      {
        "properties": {
          "distance": "6",
          "persistent": "true"
        },
        "id": 196
      },
      {
        "properties": {
          "distance": "6",
          "persistent": "false"
        },
        "id": 197
      },
      {
        "properties": {
          "distance": "7",
          "persistent": "true"
        },
        "id": 198
      },
      {
        "properties": {
          "distance": "7",
          "persistent": "false"
        },
        "id": 199,
        "default": true
      }
    ]
  },
  "minecraft:acacia_leaves": {
    "properties": {
      "distance": [
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7"
      ],
      "persistent": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "distance": "1",
          "persistent": "true"
        },
        "id": 200
      },
      {
        "properties": {
          "distance": "1",
          "persistent": "false"
        },
        "id": 201
      },
      {
        "properties": {
          "distance": "2",
          "persistent": "true"
        },
        "id": 202
      },
      {
        "properties": {
          "distance": "2",
          "persistent": "false"
        },
        "id": 203
      },
      {
        "properties": {
          "distance": "3",
          "persistent": "true"
        },
        "id": 204
      },
      {
        "properties": {
          "distance": "3",
          "persistent": "false"
        },
        "id": 205
      },
      {
        "properties": {
          "distance": "4",
          "persistent": "true"
        },
        "id": 206
      },
      {
        "properties": {
          "distance": "4",
          "persistent": "false"
        },
        "id": 207
      },
      {
        "properties": {
          "distance": "5",
          "persistent": "true"
        },
        "id": 208
      },
      {
        "properties": {
          "distance": "5",
          "persistent": "false"
        },
        "id": 209
      },
      {
        "properties": {
          "distance": "6",
          "persistent": "true"
        },
        "id": 210
      },
      {
        "properties": {
          "distance": "6",
          "persistent": "false"
        },
        "id": 211
      },
      {
        "properties": {
          "distance": "7",
          "persistent": "true"
        },
        "id": 212
      },
      {
        "properties": {
          "distance": "7",
          "persistent": "false"
        },
        "id": 213,
        "default": true
      }
    ]
  },
  "minecraft:dark_oak_leaves": {
    "properties": {
      "distance": [
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7"
      ],
      "persistent": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "distance": "1",
          "persistent": "true"
        },
        "id": 214
      },
      {
        "properties": {
          "distance": "1",
          "persistent": "false"
        },
        "id": 215
      },
      {
        "properties": {
          "distance": "2",
          "persistent": "true"
        },
        "id": 216
      },
      {
        "properties": {
          "distance": "2",
          "persistent": "false"
        },
        "id": 217
      },
      {
        "properties": {
          "distance": "3",
          "persistent": "true"
        },
        "id": 218
      },
      {
        "properties": {
          "distance": "3",
          "persistent": "false"
        },
        "id": 219
      },
      {
        "properties": {
          "distance": "4",
          "persistent": "true"
        },
        "id": 220
      },
      {
        "properties": {
          "distance": "4",
          "persistent": "false"
        },
        "id": 221
      },
      {
        "properties": {
          "distance": "5",
          "persistent": "true"
        },
        "id": 222
      },
      {
        "properties": {
          "distance": "5",
          "persistent": "false"
        },
        "id": 223
      },
      {
        "properties": {
          "distance": "6",
          "persistent": "true"
        },
        "id": 224
      },
      {
        "properties": {
          "distance": "6",
          "persistent": "false"
        },
        "id": 225
      },
      {
        "properties": {
          "distance": "7",
          "persistent": "true"
        },
        "id": 226
      },
      {
        "properties": {
          "distance": "7",
          "persistent": "false"
        },
        "id": 227,
        "default": true
      }
    ]
  },
  "minecraft:sponge": {
    "states": [
      {
        "id": 228,
        "default": true
      }
    ]
  },
  "minecraft:wet_sponge": {
    "states": [
      {
        "id": 229,
        "default": true
      }
    ]
  },
  "minecraft:glass": {
    "states": [
      {
        "id": 230,
        "default": true
      }
    ]
  },
  "minecraft:lapis_ore": {
    "states": [
      {
        "id": 231,
        "default": true
      }
    ]
  },
  "minecraft:lapis_block": {
    "states": [
      {
        "id": 232,
        "default": true
      }
    ]
  },
  "minecraft:dispenser": {
    "properties": {
      "facing": [
        "north",
        "east",
        "south",
        "west",
        "up",
        "down"
      ],
      "triggered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "triggered": "true"
        },
        "id": 233
      },
      {
        "properties": {
          "facing": "north",
          "triggered": "false"
        },
        "id": 234,
        "default": true
      },
      {
        "properties": {
          "facing": "east",
          "triggered": "true"
        },
        "id": 235
      },
      {
        "properties": {
          "facing": "east",
          "triggered": "false"
        },
        "id": 236
      },
      {
        "properties": {
          "facing": "south",
          "triggered": "true"
        },
        "id": 237
      },
      {
        "properties": {
          "facing": "south",
          "triggered": "false"
        },
        "id": 238
      },
      {
        "properties": {
          "facing": "west",
          "triggered": "true"
        },
        "id": 239
      },
      {
        "properties": {
          "facing": "west",
          "triggered": "false"
        },
        "id": 240
      },
      {
        "properties": {
          "facing": "up",
          "triggered": "true"
        },
        "id": 241
      },
      {
        "properties": {
          "facing": "up",
          "triggered": "false"
        },
        "id": 242
      },
      {
        "properties": {
          "facing": "down",
          "triggered": "true"
        },
        "id": 243
      },
      {
        "properties": {
          "facing": "down",
          "triggered": "false"
        },
        "id": 244
      }
    ]
  },
  "minecraft:sandstone": {
    "states": [
      {
        "id": 245,
        "default": true
      }
    ]
  },
  "minecraft:chiseled_sandstone": {
    "states": [
      {
        "id": 246,
        "default": true
      }
    ]
  },
  "minecraft:cut_sandstone": {
    "states": [
      {
        "id": 247,
        "default": true
      }
    ]
  },
  "minecraft:note_block": {
    "properties": {
      "instrument": [
        "harp",
        "basedrum",
        "snare",
        "hat",
        "bass",
        "flute",
        "bell",
        "guitar",
        "chime",
        "xylophone"
      ],
      "note": [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "10",
        "11",
        "12",
        "13",
        "14",
        "15",
        "16",
        "17",
        "18",
        "19",
        "20",
        "21",
        "22",
        "23",
        "24"
      ],
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "instrument": "harp",
          "note": "0",
          "powered": "true"
        },
        "id": 248
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "0",
          "powered": "false"
        },
        "id": 249,
        "default": true
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "1",
          "powered": "true"
        },
        "id": 250
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "1",
          "powered": "false"
        },
        "id": 251
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "2",
          "powered": "true"
        },
        "id": 252
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "2",
          "powered": "false"
        },
        "id": 253
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "3",
          "powered": "true"
        },
        "id": 254
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "3",
          "powered": "false"
        },
        "id": 255
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "4",
          "powered": "true"
        },
        "id": 256
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "4",
          "powered": "false"
        },
        "id": 257
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "5",
          "powered": "true"
        },
        "id": 258
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "5",
          "powered": "false"
        },
        "id": 259
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "6",
          "powered": "true"
        },
        "id": 260
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "6",
          "powered": "false"
        },
        "id": 261
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "7",
          "powered": "true"
        },
        "id": 262
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "7",
          "powered": "false"
        },
        "id": 263
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "8",
          "powered": "true"
        },
        "id": 264
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "8",
          "powered": "false"
        },
        "id": 265
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "9",
          "powered": "true"
        },
        "id": 266
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "9",
          "powered": "false"
        },
        "id": 267
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "10",
          "powered": "true"
        },
        "id": 268
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "10",
          "powered": "false"
        },
        "id": 269
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "11",
          "powered": "true"
        },
        "id": 270
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "11",
          "powered": "false"
        },
        "id": 271
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "12",
          "powered": "true"
        },
        "id": 272
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "12",
          "powered": "false"
        },
        "id": 273
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "13",
          "powered": "true"
        },
        "id": 274
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "13",
          "powered": "false"
        },
        "id": 275
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "14",
          "powered": "true"
        },
        "id": 276
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "14",
          "powered": "false"
        },
        "id": 277
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "15",
          "powered": "true"
        },
        "id": 278
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "15",
          "powered": "false"
        },
        "id": 279
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "16",
          "powered": "true"
        },
        "id": 280
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "16",
          "powered": "false"
        },
        "id": 281
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "17",
          "powered": "true"
        },
        "id": 282
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "17",
          "powered": "false"
        },
        "id": 283
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "18",
          "powered": "true"
        },
        "id": 284
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "18",
          "powered": "false"
        },
        "id": 285
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "19",
          "powered": "true"
        },
        "id": 286
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "19",
          "powered": "false"
        },
        "id": 287
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "20",
          "powered": "true"
        },
        "id": 288
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "20",
          "powered": "false"
        },
        "id": 289
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "21",
          "powered": "true"
        },
        "id": 290
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "21",
          "powered": "false"
        },
        "id": 291
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "22",
          "powered": "true"
        },
        "id": 292
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "22",
          "powered": "false"
        },
        "id": 293
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "23",
          "powered": "true"
        },
        "id": 294
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "23",
          "powered": "false"
        },
        "id": 295
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "24",
          "powered": "true"
        },
        "id": 296
      },
      {
        "properties": {
          "instrument": "harp",
          "note": "24",
          "powered": "false"
        },
        "id": 297
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "0",
          "powered": "true"
        },
        "id": 298
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "0",
          "powered": "false"
        },
        "id": 299
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "1",
          "powered": "true"
        },
        "id": 300
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "1",
          "powered": "false"
        },
        "id": 301
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "2",
          "powered": "true"
        },
        "id": 302
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "2",
          "powered": "false"
        },
        "id": 303
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "3",
          "powered": "true"
        },
        "id": 304
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "3",
          "powered": "false"
        },
        "id": 305
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "4",
          "powered": "true"
        },
        "id": 306
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "4",
          "powered": "false"
        },
        "id": 307
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "5",
          "powered": "true"
        },
        "id": 308
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "5",
          "powered": "false"
        },
        "id": 309
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "6",
          "powered": "true"
        },
        "id": 310
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "6",
          "powered": "false"
        },
        "id": 311
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "7",
          "powered": "true"
        },
        "id": 312
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "7",
          "powered": "false"
        },
        "id": 313
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "8",
          "powered": "true"
        },
        "id": 314
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "8",
          "powered": "false"
        },
        "id": 315
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "9",
          "powered": "true"
        },
        "id": 316
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "9",
          "powered": "false"
        },
        "id": 317
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "10",
          "powered": "true"
        },
        "id": 318
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "10",
          "powered": "false"
        },
        "id": 319
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "11",
          "powered": "true"
        },
        "id": 320
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "11",
          "powered": "false"
        },
        "id": 321
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "12",
          "powered": "true"
        },
        "id": 322
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "12",
          "powered": "false"
        },
        "id": 323
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "13",
          "powered": "true"
        },
        "id": 324
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "13",
          "powered": "false"
        },
        "id": 325
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "14",
          "powered": "true"
        },
        "id": 326
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "14",
          "powered": "false"
        },
        "id": 327
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "15",
          "powered": "true"
        },
        "id": 328
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "15",
          "powered": "false"
        },
        "id": 329
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "16",
          "powered": "true"
        },
        "id": 330
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "16",
          "powered": "false"
        },
        "id": 331
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "17",
          "powered": "true"
        },
        "id": 332
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "17",
          "powered": "false"
        },
        "id": 333
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "18",
          "powered": "true"
        },
        "id": 334
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "18",
          "powered": "false"
        },
        "id": 335
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "19",
          "powered": "true"
        },
        "id": 336
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "19",
          "powered": "false"
        },
        "id": 337
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "20",
          "powered": "true"
        },
        "id": 338
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "20",
          "powered": "false"
        },
        "id": 339
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "21",
          "powered": "true"
        },
        "id": 340
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "21",
          "powered": "false"
        },
        "id": 341
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "22",
          "powered": "true"
        },
        "id": 342
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "22",
          "powered": "false"
        },
        "id": 343
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "23",
          "powered": "true"
        },
        "id": 344
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "23",
          "powered": "false"
        },
        "id": 345
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "24",
          "powered": "true"
        },
        "id": 346
      },
      {
        "properties": {
          "instrument": "basedrum",
          "note": "24",
          "powered": "false"
        },
        "id": 347
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "0",
          "powered": "true"
        },
        "id": 348
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "0",
          "powered": "false"
        },
        "id": 349
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "1",
          "powered": "true"
        },
        "id": 350
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "1",
          "powered": "false"
        },
        "id": 351
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "2",
          "powered": "true"
        },
        "id": 352
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "2",
          "powered": "false"
        },
        "id": 353
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "3",
          "powered": "true"
        },
        "id": 354
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "3",
          "powered": "false"
        },
        "id": 355
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "4",
          "powered": "true"
        },
        "id": 356
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "4",
          "powered": "false"
        },
        "id": 357
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "5",
          "powered": "true"
        },
        "id": 358
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "5",
          "powered": "false"
        },
        "id": 359
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "6",
          "powered": "true"
        },
        "id": 360
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "6",
          "powered": "false"
        },
        "id": 361
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "7",
          "powered": "true"
        },
        "id": 362
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "7",
          "powered": "false"
        },
        "id": 363
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "8",
          "powered": "true"
        },
        "id": 364
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "8",
          "powered": "false"
        },
        "id": 365
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "9",
          "powered": "true"
        },
        "id": 366
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "9",
          "powered": "false"
        },
        "id": 367
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "10",
          "powered": "true"
        },
        "id": 368
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "10",
          "powered": "false"
        },
        "id": 369
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "11",
          "powered": "true"
        },
        "id": 370
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "11",
          "powered": "false"
        },
        "id": 371
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "12",
          "powered": "true"
        },
        "id": 372
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "12",
          "powered": "false"
        },
        "id": 373
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "13",
          "powered": "true"
        },
        "id": 374
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "13",
          "powered": "false"
        },
        "id": 375
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "14",
          "powered": "true"
        },
        "id": 376
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "14",
          "powered": "false"
        },
        "id": 377
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "15",
          "powered": "true"
        },
        "id": 378
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "15",
          "powered": "false"
        },
        "id": 379
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "16",
          "powered": "true"
        },
        "id": 380
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "16",
          "powered": "false"
        },
        "id": 381
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "17",
          "powered": "true"
        },
        "id": 382
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "17",
          "powered": "false"
        },
        "id": 383
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "18",
          "powered": "true"
        },
        "id": 384
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "18",
          "powered": "false"
        },
        "id": 385
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "19",
          "powered": "true"
        },
        "id": 386
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "19",
          "powered": "false"
        },
        "id": 387
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "20",
          "powered": "true"
        },
        "id": 388
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "20",
          "powered": "false"
        },
        "id": 389
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "21",
          "powered": "true"
        },
        "id": 390
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "21",
          "powered": "false"
        },
        "id": 391
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "22",
          "powered": "true"
        },
        "id": 392
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "22",
          "powered": "false"
        },
        "id": 393
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "23",
          "powered": "true"
        },
        "id": 394
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "23",
          "powered": "false"
        },
        "id": 395
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "24",
          "powered": "true"
        },
        "id": 396
      },
      {
        "properties": {
          "instrument": "snare",
          "note": "24",
          "powered": "false"
        },
        "id": 397
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "0",
          "powered": "true"
        },
        "id": 398
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "0",
          "powered": "false"
        },
        "id": 399
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "1",
          "powered": "true"
        },
        "id": 400
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "1",
          "powered": "false"
        },
        "id": 401
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "2",
          "powered": "true"
        },
        "id": 402
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "2",
          "powered": "false"
        },
        "id": 403
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "3",
          "powered": "true"
        },
        "id": 404
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "3",
          "powered": "false"
        },
        "id": 405
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "4",
          "powered": "true"
        },
        "id": 406
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "4",
          "powered": "false"
        },
        "id": 407
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "5",
          "powered": "true"
        },
        "id": 408
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "5",
          "powered": "false"
        },
        "id": 409
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "6",
          "powered": "true"
        },
        "id": 410
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "6",
          "powered": "false"
        },
        "id": 411
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "7",
          "powered": "true"
        },
        "id": 412
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "7",
          "powered": "false"
        },
        "id": 413
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "8",
          "powered": "true"
        },
        "id": 414
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "8",
          "powered": "false"
        },
        "id": 415
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "9",
          "powered": "true"
        },
        "id": 416
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "9",
          "powered": "false"
        },
        "id": 417
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "10",
          "powered": "true"
        },
        "id": 418
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "10",
          "powered": "false"
        },
        "id": 419
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "11",
          "powered": "true"
        },
        "id": 420
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "11",
          "powered": "false"
        },
        "id": 421
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "12",
          "powered": "true"
        },
        "id": 422
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "12",
          "powered": "false"
        },
        "id": 423
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "13",
          "powered": "true"
        },
        "id": 424
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "13",
          "powered": "false"
        },
        "id": 425
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "14",
          "powered": "true"
        },
        "id": 426
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "14",
          "powered": "false"
        },
        "id": 427
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "15",
          "powered": "true"
        },
        "id": 428
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "15",
          "powered": "false"
        },
        "id": 429
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "16",
          "powered": "true"
        },
        "id": 430
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "16",
          "powered": "false"
        },
        "id": 431
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "17",
          "powered": "true"
        },
        "id": 432
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "17",
          "powered": "false"
        },
        "id": 433
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "18",
          "powered": "true"
        },
        "id": 434
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "18",
          "powered": "false"
        },
        "id": 435
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "19",
          "powered": "true"
        },
        "id": 436
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "19",
          "powered": "false"
        },
        "id": 437
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "20",
          "powered": "true"
        },
        "id": 438
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "20",
          "powered": "false"
        },
        "id": 439
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "21",
          "powered": "true"
        },
        "id": 440
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "21",
          "powered": "false"
        },
        "id": 441
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "22",
          "powered": "true"
        },
        "id": 442
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "22",
          "powered": "false"
        },
        "id": 443
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "23",
          "powered": "true"
        },
        "id": 444
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "23",
          "powered": "false"
        },
        "id": 445
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "24",
          "powered": "true"
        },
        "id": 446
      },
      {
        "properties": {
          "instrument": "hat",
          "note": "24",
          "powered": "false"
        },
        "id": 447
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "0",
          "powered": "true"
        },
        "id": 448
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "0",
          "powered": "false"
        },
        "id": 449
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "1",
          "powered": "true"
        },
        "id": 450
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "1",
          "powered": "false"
        },
        "id": 451
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "2",
          "powered": "true"
        },
        "id": 452
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "2",
          "powered": "false"
        },
        "id": 453
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "3",
          "powered": "true"
        },
        "id": 454
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "3",
          "powered": "false"
        },
        "id": 455
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "4",
          "powered": "true"
        },
        "id": 456
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "4",
          "powered": "false"
        },
        "id": 457
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "5",
          "powered": "true"
        },
        "id": 458
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "5",
          "powered": "false"
        },
        "id": 459
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "6",
          "powered": "true"
        },
        "id": 460
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "6",
          "powered": "false"
        },
        "id": 461
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "7",
          "powered": "true"
        },
        "id": 462
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "7",
          "powered": "false"
        },
        "id": 463
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "8",
          "powered": "true"
        },
        "id": 464
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "8",
          "powered": "false"
        },
        "id": 465
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "9",
          "powered": "true"
        },
        "id": 466
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "9",
          "powered": "false"
        },
        "id": 467
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "10",
          "powered": "true"
        },
        "id": 468
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "10",
          "powered": "false"
        },
        "id": 469
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "11",
          "powered": "true"
        },
        "id": 470
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "11",
          "powered": "false"
        },
        "id": 471
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "12",
          "powered": "true"
        },
        "id": 472
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "12",
          "powered": "false"
        },
        "id": 473
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "13",
          "powered": "true"
        },
        "id": 474
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "13",
          "powered": "false"
        },
        "id": 475
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "14",
          "powered": "true"
        },
        "id": 476
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "14",
          "powered": "false"
        },
        "id": 477
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "15",
          "powered": "true"
        },
        "id": 478
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "15",
          "powered": "false"
        },
        "id": 479
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "16",
          "powered": "true"
        },
        "id": 480
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "16",
          "powered": "false"
        },
        "id": 481
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "17",
          "powered": "true"
        },
        "id": 482
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "17",
          "powered": "false"
        },
        "id": 483
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "18",
          "powered": "true"
        },
        "id": 484
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "18",
          "powered": "false"
        },
        "id": 485
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "19",
          "powered": "true"
        },
        "id": 486
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "19",
          "powered": "false"
        },
        "id": 487
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "20",
          "powered": "true"
        },
        "id": 488
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "20",
          "powered": "false"
        },
        "id": 489
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "21",
          "powered": "true"
        },
        "id": 490
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "21",
          "powered": "false"
        },
        "id": 491
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "22",
          "powered": "true"
        },
        "id": 492
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "22",
          "powered": "false"
        },
        "id": 493
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "23",
          "powered": "true"
        },
        "id": 494
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "23",
          "powered": "false"
        },
        "id": 495
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "24",
          "powered": "true"
        },
        "id": 496
      },
      {
        "properties": {
          "instrument": "bass",
          "note": "24",
          "powered": "false"
        },
        "id": 497
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "0",
          "powered": "true"
        },
        "id": 498
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "0",
          "powered": "false"
        },
        "id": 499
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "1",
          "powered": "true"
        },
        "id": 500
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "1",
          "powered": "false"
        },
        "id": 501
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "2",
          "powered": "true"
        },
        "id": 502
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "2",
          "powered": "false"
        },
        "id": 503
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "3",
          "powered": "true"
        },
        "id": 504
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "3",
          "powered": "false"
        },
        "id": 505
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "4",
          "powered": "true"
        },
        "id": 506
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "4",
          "powered": "false"
        },
        "id": 507
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "5",
          "powered": "true"
        },
        "id": 508
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "5",
          "powered": "false"
        },
        "id": 509
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "6",
          "powered": "true"
        },
        "id": 510
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "6",
          "powered": "false"
        },
        "id": 511
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "7",
          "powered": "true"
        },
        "id": 512
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "7",
          "powered": "false"
        },
        "id": 513
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "8",
          "powered": "true"
        },
        "id": 514
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "8",
          "powered": "false"
        },
        "id": 515
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "9",
          "powered": "true"
        },
        "id": 516
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "9",
          "powered": "false"
        },
        "id": 517
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "10",
          "powered": "true"
        },
        "id": 518
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "10",
          "powered": "false"
        },
        "id": 519
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "11",
          "powered": "true"
        },
        "id": 520
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "11",
          "powered": "false"
        },
        "id": 521
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "12",
          "powered": "true"
        },
        "id": 522
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "12",
          "powered": "false"
        },
        "id": 523
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "13",
          "powered": "true"
        },
        "id": 524
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "13",
          "powered": "false"
        },
        "id": 525
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "14",
          "powered": "true"
        },
        "id": 526
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "14",
          "powered": "false"
        },
        "id": 527
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "15",
          "powered": "true"
        },
        "id": 528
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "15",
          "powered": "false"
        },
        "id": 529
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "16",
          "powered": "true"
        },
        "id": 530
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "16",
          "powered": "false"
        },
        "id": 531
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "17",
          "powered": "true"
        },
        "id": 532
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "17",
          "powered": "false"
        },
        "id": 533
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "18",
          "powered": "true"
        },
        "id": 534
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "18",
          "powered": "false"
        },
        "id": 535
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "19",
          "powered": "true"
        },
        "id": 536
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "19",
          "powered": "false"
        },
        "id": 537
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "20",
          "powered": "true"
        },
        "id": 538
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "20",
          "powered": "false"
        },
        "id": 539
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "21",
          "powered": "true"
        },
        "id": 540
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "21",
          "powered": "false"
        },
        "id": 541
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "22",
          "powered": "true"
        },
        "id": 542
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "22",
          "powered": "false"
        },
        "id": 543
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "23",
          "powered": "true"
        },
        "id": 544
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "23",
          "powered": "false"
        },
        "id": 545
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "24",
          "powered": "true"
        },
        "id": 546
      },
      {
        "properties": {
          "instrument": "flute",
          "note": "24",
          "powered": "false"
        },
        "id": 547
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "0",
          "powered": "true"
        },
        "id": 548
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "0",
          "powered": "false"
        },
        "id": 549
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "1",
          "powered": "true"
        },
        "id": 550
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "1",
          "powered": "false"
        },
        "id": 551
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "2",
          "powered": "true"
        },
        "id": 552
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "2",
          "powered": "false"
        },
        "id": 553
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "3",
          "powered": "true"
        },
        "id": 554
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "3",
          "powered": "false"
        },
        "id": 555
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "4",
          "powered": "true"
        },
        "id": 556
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "4",
          "powered": "false"
        },
        "id": 557
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "5",
          "powered": "true"
        },
        "id": 558
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "5",
          "powered": "false"
        },
        "id": 559
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "6",
          "powered": "true"
        },
        "id": 560
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "6",
          "powered": "false"
        },
        "id": 561
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "7",
          "powered": "true"
        },
        "id": 562
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "7",
          "powered": "false"
        },
        "id": 563
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "8",
          "powered": "true"
        },
        "id": 564
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "8",
          "powered": "false"
        },
        "id": 565
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "9",
          "powered": "true"
        },
        "id": 566
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "9",
          "powered": "false"
        },
        "id": 567
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "10",
          "powered": "true"
        },
        "id": 568
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "10",
          "powered": "false"
        },
        "id": 569
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "11",
          "powered": "true"
        },
        "id": 570
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "11",
          "powered": "false"
        },
        "id": 571
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "12",
          "powered": "true"
        },
        "id": 572
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "12",
          "powered": "false"
        },
        "id": 573
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "13",
          "powered": "true"
        },
        "id": 574
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "13",
          "powered": "false"
        },
        "id": 575
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "14",
          "powered": "true"
        },
        "id": 576
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "14",
          "powered": "false"
        },
        "id": 577
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "15",
          "powered": "true"
        },
        "id": 578
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "15",
          "powered": "false"
        },
        "id": 579
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "16",
          "powered": "true"
        },
        "id": 580
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "16",
          "powered": "false"
        },
        "id": 581
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "17",
          "powered": "true"
        },
        "id": 582
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "17",
          "powered": "false"
        },
        "id": 583
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "18",
          "powered": "true"
        },
        "id": 584
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "18",
          "powered": "false"
        },
        "id": 585
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "19",
          "powered": "true"
        },
        "id": 586
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "19",
          "powered": "false"
        },
        "id": 587
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "20",
          "powered": "true"
        },
        "id": 588
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "20",
          "powered": "false"
        },
        "id": 589
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "21",
          "powered": "true"
        },
        "id": 590
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "21",
          "powered": "false"
        },
        "id": 591
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "22",
          "powered": "true"
        },
        "id": 592
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "22",
          "powered": "false"
        },
        "id": 593
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "23",
          "powered": "true"
        },
        "id": 594
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "23",
          "powered": "false"
        },
        "id": 595
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "24",
          "powered": "true"
        },
        "id": 596
      },
      {
        "properties": {
          "instrument": "bell",
          "note": "24",
          "powered": "false"
        },
        "id": 597
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "0",
          "powered": "true"
        },
        "id": 598
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "0",
          "powered": "false"
        },
        "id": 599
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "1",
          "powered": "true"
        },
        "id": 600
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "1",
          "powered": "false"
        },
        "id": 601
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "2",
          "powered": "true"
        },
        "id": 602
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "2",
          "powered": "false"
        },
        "id": 603
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "3",
          "powered": "true"
        },
        "id": 604
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "3",
          "powered": "false"
        },
        "id": 605
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "4",
          "powered": "true"
        },
        "id": 606
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "4",
          "powered": "false"
        },
        "id": 607
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "5",
          "powered": "true"
        },
        "id": 608
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "5",
          "powered": "false"
        },
        "id": 609
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "6",
          "powered": "true"
        },
        "id": 610
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "6",
          "powered": "false"
        },
        "id": 611
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "7",
          "powered": "true"
        },
        "id": 612
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "7",
          "powered": "false"
        },
        "id": 613
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "8",
          "powered": "true"
        },
        "id": 614
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "8",
          "powered": "false"
        },
        "id": 615
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "9",
          "powered": "true"
        },
        "id": 616
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "9",
          "powered": "false"
        },
        "id": 617
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "10",
          "powered": "true"
        },
        "id": 618
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "10",
          "powered": "false"
        },
        "id": 619
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "11",
          "powered": "true"
        },
        "id": 620
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "11",
          "powered": "false"
        },
        "id": 621
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "12",
          "powered": "true"
        },
        "id": 622
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "12",
          "powered": "false"
        },
        "id": 623
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "13",
          "powered": "true"
        },
        "id": 624
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "13",
          "powered": "false"
        },
        "id": 625
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "14",
          "powered": "true"
        },
        "id": 626
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "14",
          "powered": "false"
        },
        "id": 627
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "15",
          "powered": "true"
        },
        "id": 628
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "15",
          "powered": "false"
        },
        "id": 629
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "16",
          "powered": "true"
        },
        "id": 630
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "16",
          "powered": "false"
        },
        "id": 631
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "17",
          "powered": "true"
        },
        "id": 632
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "17",
          "powered": "false"
        },
        "id": 633
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "18",
          "powered": "true"
        },
        "id": 634
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "18",
          "powered": "false"
        },
        "id": 635
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "19",
          "powered": "true"
        },
        "id": 636
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "19",
          "powered": "false"
        },
        "id": 637
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "20",
          "powered": "true"
        },
        "id": 638
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "20",
          "powered": "false"
        },
        "id": 639
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "21",
          "powered": "true"
        },
        "id": 640
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "21",
          "powered": "false"
        },
        "id": 641
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "22",
          "powered": "true"
        },
        "id": 642
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "22",
          "powered": "false"
        },
        "id": 643
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "23",
          "powered": "true"
        },
        "id": 644
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "23",
          "powered": "false"
        },
        "id": 645
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "24",
          "powered": "true"
        },
        "id": 646
      },
      {
        "properties": {
          "instrument": "guitar",
          "note": "24",
          "powered": "false"
        },
        "id": 647
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "0",
          "powered": "true"
        },
        "id": 648
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "0",
          "powered": "false"
        },
        "id": 649
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "1",
          "powered": "true"
        },
        "id": 650
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "1",
          "powered": "false"
        },
        "id": 651
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "2",
          "powered": "true"
        },
        "id": 652
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "2",
          "powered": "false"
        },
        "id": 653
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "3",
          "powered": "true"
        },
        "id": 654
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "3",
          "powered": "false"
        },
        "id": 655
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "4",
          "powered": "true"
        },
        "id": 656
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "4",
          "powered": "false"
        },
        "id": 657
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "5",
          "powered": "true"
        },
        "id": 658
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "5",
          "powered": "false"
        },
        "id": 659
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "6",
          "powered": "true"
        },
        "id": 660
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "6",
          "powered": "false"
        },
        "id": 661
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "7",
          "powered": "true"
        },
        "id": 662
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "7",
          "powered": "false"
        },
        "id": 663
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "8",
          "powered": "true"
        },
        "id": 664
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "8",
          "powered": "false"
        },
        "id": 665
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "9",
          "powered": "true"
        },
        "id": 666
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "9",
          "powered": "false"
        },
        "id": 667
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "10",
          "powered": "true"
        },
        "id": 668
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "10",
          "powered": "false"
        },
        "id": 669
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "11",
          "powered": "true"
        },
        "id": 670
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "11",
          "powered": "false"
        },
        "id": 671
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "12",
          "powered": "true"
        },
        "id": 672
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "12",
          "powered": "false"
        },
        "id": 673
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "13",
          "powered": "true"
        },
        "id": 674
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "13",
          "powered": "false"
        },
        "id": 675
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "14",
          "powered": "true"
        },
        "id": 676
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "14",
          "powered": "false"
        },
        "id": 677
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "15",
          "powered": "true"
        },
        "id": 678
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "15",
          "powered": "false"
        },
        "id": 679
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "16",
          "powered": "true"
        },
        "id": 680
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "16",
          "powered": "false"
        },
        "id": 681
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "17",
          "powered": "true"
        },
        "id": 682
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "17",
          "powered": "false"
        },
        "id": 683
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "18",
          "powered": "true"
        },
        "id": 684
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "18",
          "powered": "false"
        },
        "id": 685
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "19",
          "powered": "true"
        },
        "id": 686
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "19",
          "powered": "false"
        },
        "id": 687
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "20",
          "powered": "true"
        },
        "id": 688
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "20",
          "powered": "false"
        },
        "id": 689
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "21",
          "powered": "true"
        },
        "id": 690
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "21",
          "powered": "false"
        },
        "id": 691
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "22",
          "powered": "true"
        },
        "id": 692
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "22",
          "powered": "false"
        },
        "id": 693
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "23",
          "powered": "true"
        },
        "id": 694
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "23",
          "powered": "false"
        },
        "id": 695
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "24",
          "powered": "true"
        },
        "id": 696
      },
      {
        "properties": {
          "instrument": "chime",
          "note": "24",
          "powered": "false"
        },
        "id": 697
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "0",
          "powered": "true"
        },
        "id": 698
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "0",
          "powered": "false"
        },
        "id": 699
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "1",
          "powered": "true"
        },
        "id": 700
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "1",
          "powered": "false"
        },
        "id": 701
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "2",
          "powered": "true"
        },
        "id": 702
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "2",
          "powered": "false"
        },
        "id": 703
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "3",
          "powered": "true"
        },
        "id": 704
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "3",
          "powered": "false"
        },
        "id": 705
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "4",
          "powered": "true"
        },
        "id": 706
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "4",
          "powered": "false"
        },
        "id": 707
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "5",
          "powered": "true"
        },
        "id": 708
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "5",
          "powered": "false"
        },
        "id": 709
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "6",
          "powered": "true"
        },
        "id": 710
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "6",
          "powered": "false"
        },
        "id": 711
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "7",
          "powered": "true"
        },
        "id": 712
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "7",
          "powered": "false"
        },
        "id": 713
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "8",
          "powered": "true"
        },
        "id": 714
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "8",
          "powered": "false"
        },
        "id": 715
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "9",
          "powered": "true"
        },
        "id": 716
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "9",
          "powered": "false"
        },
        "id": 717
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "10",
          "powered": "true"
        },
        "id": 718
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "10",
          "powered": "false"
        },
        "id": 719
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "11",
          "powered": "true"
        },
        "id": 720
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "11",
          "powered": "false"
        },
        "id": 721
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "12",
          "powered": "true"
        },
        "id": 722
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "12",
          "powered": "false"
        },
        "id": 723
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "13",
          "powered": "true"
        },
        "id": 724
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "13",
          "powered": "false"
        },
        "id": 725
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "14",
          "powered": "true"
        },
        "id": 726
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "14",
          "powered": "false"
        },
        "id": 727
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "15",
          "powered": "true"
        },
        "id": 728
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "15",
          "powered": "false"
        },
        "id": 729
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "16",
          "powered": "true"
        },
        "id": 730
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "16",
          "powered": "false"
        },
        "id": 731
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "17",
          "powered": "true"
        },
        "id": 732
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "17",
          "powered": "false"
        },
        "id": 733
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "18",
          "powered": "true"
        },
        "id": 734
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "18",
          "powered": "false"
        },
        "id": 735
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "19",
          "powered": "true"
        },
        "id": 736
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "19",
          "powered": "false"
        },
        "id": 737
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "20",
          "powered": "true"
        },
        "id": 738
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "20",
          "powered": "false"
        },
        "id": 739
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "21",
          "powered": "true"
        },
        "id": 740
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "21",
          "powered": "false"
        },
        "id": 741
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "22",
          "powered": "true"
        },
        "id": 742
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "22",
          "powered": "false"
        },
        "id": 743
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "23",
          "powered": "true"
        },
        "id": 744
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "23",
          "powered": "false"
        },
        "id": 745
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "24",
          "powered": "true"
        },
        "id": 746
      },
      {
        "properties": {
          "instrument": "xylophone",
          "note": "24",
          "powered": "false"
        },
        "id": 747
      }
    ]
  },
  "minecraft:white_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 748
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 749
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 750
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 751,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 752
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 753
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 754
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 755
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 756
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 757
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 758
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 759
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 760
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 761
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 762
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 763
      }
    ]
  },
  "minecraft:orange_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 764
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 765
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 766
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 767,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 768
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 769
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 770
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 771
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 772
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 773
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 774
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 775
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 776
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 777
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 778
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 779
      }
    ]
  },
  "minecraft:magenta_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 780
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 781
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 782
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 783,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 784
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 785
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 786
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 787
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 788
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 789
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 790
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 791
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 792
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 793
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 794
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 795
      }
    ]
  },
  "minecraft:light_blue_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 796
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 797
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 798
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 799,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 800
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 801
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 802
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 803
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 804
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 805
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 806
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 807
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 808
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 809
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 810
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 811
      }
    ]
  },
  "minecraft:yellow_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 812
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 813
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 814
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 815,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 816
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 817
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 818
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 819
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 820
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 821
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 822
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 823
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 824
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 825
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 826
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 827
      }
    ]
  },
  "minecraft:lime_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 828
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 829
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 830
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 831,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 832
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 833
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 834
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 835
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 836
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 837
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 838
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 839
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 840
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 841
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 842
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 843
      }
    ]
  },
  "minecraft:pink_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 844
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 845
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 846
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 847,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 848
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 849
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 850
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 851
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 852
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 853
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 854
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 855
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 856
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 857
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 858
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 859
      }
    ]
  },
  "minecraft:gray_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 860
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 861
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 862
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 863,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 864
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 865
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 866
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 867
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 868
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 869
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 870
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 871
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 872
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 873
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 874
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 875
      }
    ]
  },
  "minecraft:light_gray_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 876
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 877
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 878
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 879,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 880
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 881
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 882
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 883
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 884
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 885
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 886
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 887
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 888
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 889
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 890
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 891
      }
    ]
  },
  "minecraft:cyan_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 892
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 893
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 894
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 895,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 896
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 897
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 898
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 899
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 900
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 901
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 902
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 903
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 904
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 905
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 906
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 907
      }
    ]
  },
  "minecraft:purple_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 908
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 909
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 910
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 911,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 912
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 913
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 914
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 915
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 916
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 917
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 918
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 919
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 920
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 921
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 922
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 923
      }
    ]
  },
  "minecraft:blue_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 924
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 925
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 926
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 927,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 928
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 929
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 930
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 931
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 932
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 933
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 934
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 935
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 936
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 937
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 938
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 939
      }
    ]
  },
  "minecraft:brown_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 940
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 941
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 942
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 943,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 944
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 945
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 946
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 947
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 948
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 949
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 950
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 951
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 952
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 953
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 954
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 955
      }
    ]
  },
  "minecraft:green_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 956
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 957
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 958
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 959,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 960
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 961
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 962
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 963
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 964
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 965
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 966
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 967
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 968
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 969
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 970
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 971
      }
    ]
  },
  "minecraft:red_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 972
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 973
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 974
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 975,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 976
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 977
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 978
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 979
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 980
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 981
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 982
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 983
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 984
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 985
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 986
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 987
      }
    ]
  },
  "minecraft:black_bed": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "occupied": [
        "true",
        "false"
      ],
      "part": [
        "head",
        "foot"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "head"
        },
        "id": 988
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "true",
          "part": "foot"
        },
        "id": 989
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "head"
        },
        "id": 990
      },
      {
        "properties": {
          "facing": "north",
          "occupied": "false",
          "part": "foot"
        },
        "id": 991,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "head"
        },
        "id": 992
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "true",
          "part": "foot"
        },
        "id": 993
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "head"
        },
        "id": 994
      },
      {
        "properties": {
          "facing": "south",
          "occupied": "false",
          "part": "foot"
        },
        "id": 995
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "head"
        },
        "id": 996
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "true",
          "part": "foot"
        },
        "id": 997
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "head"
        },
        "id": 998
      },
      {
        "properties": {
          "facing": "west",
          "occupied": "false",
          "part": "foot"
        },
        "id": 999
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "head"
        },
        "id": 1000
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "true",
          "part": "foot"
        },
        "id": 1001
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "head"
        },
        "id": 1002
      },
      {
        "properties": {
          "facing": "east",
          "occupied": "false",
          "part": "foot"
        },
        "id": 1003
      }
    ]
  },
  "minecraft:powered_rail": {
    "properties": {
      "powered": [
        "true",
        "false"
      ],
      "shape": [
        "north_south",
        "east_west",
        "ascending_east",
        "ascending_west",
        "ascending_north",
        "ascending_south"
      ]
    },
    "states": [
      {
        "properties": {
          "powered": "true",
          "shape": "north_south"
        },
        "id": 1004
      },
      {
        "properties": {
          "powered": "true",
          "shape": "east_west"
        },
        "id": 1005
      },
      {
        "properties": {
          "powered": "true",
          "shape": "ascending_east"
        },
        "id": 1006
      },
      {
        "properties": {
          "powered": "true",
          "shape": "ascending_west"
        },
        "id": 1007
      },
      {
        "properties": {
          "powered": "true",
          "shape": "ascending_north"
        },
        "id": 1008
      },
      {
        "properties": {
          "powered": "true",
          "shape": "ascending_south"
        },
        "id": 1009
      },
      {
        "properties": {
          "powered": "false",
          "shape": "north_south"
        },
        "id": 1010,
        "default": true
      },
      {
        "properties": {
          "powered": "false",
          "shape": "east_west"
        },
        "id": 1011
      },
      {
        "properties": {
          "powered": "false",
          "shape": "ascending_east"
        },
        "id": 1012
      },
      {
        "properties": {
          "powered": "false",
          "shape": "ascending_west"
        },
        "id": 1013
      },
      {
        "properties": {
          "powered": "false",
          "shape": "ascending_north"
        },
        "id": 1014
      },
      {
        "properties": {
          "powered": "false",
          "shape": "ascending_south"
        },
        "id": 1015
      }
    ]
  },
  "minecraft:detector_rail": {
    "properties": {
      "powered": [
        "true",
        "false"
      ],
      "shape": [
        "north_south",
        "east_west",
        "ascending_east",
        "ascending_west",
        "ascending_north",
        "ascending_south"
      ]
    },
    "states": [
      {
        "properties": {
          "powered": "true",
          "shape": "north_south"
        },
        "id": 1016
      },
      {
        "properties": {
          "powered": "true",
          "shape": "east_west"
        },
        "id": 1017
      },
      {
        "properties": {
          "powered": "true",
          "shape": "ascending_east"
        },
        "id": 1018
      },
      {
        "properties": {
          "powered": "true",
          "shape": "ascending_west"
        },
        "id": 1019
      },
      {
        "properties": {
          "powered": "true",
          "shape": "ascending_north"
        },
        "id": 1020
      },
      {
        "properties": {
          "powered": "true",
          "shape": "ascending_south"
        },
        "id": 1021
      },
      {
        "properties": {
          "powered": "false",
          "shape": "north_south"
        },
        "id": 1022,
        "default": true
      },
      {
        "properties": {
          "powered": "false",
          "shape": "east_west"
        },
        "id": 1023
      },
      {
        "properties": {
          "powered": "false",
          "shape": "ascending_east"
        },
        "id": 1024
      },
      {
        "properties": {
          "powered": "false",
          "shape": "ascending_west"
        },
        "id": 1025
      },
      {
        "properties": {
          "powered": "false",
          "shape": "ascending_north"
        },
        "id": 1026
      },
      {
        "properties": {
          "powered": "false",
          "shape": "ascending_south"
        },
        "id": 1027
      }
    ]
  },
  "minecraft:sticky_piston": {
    "properties": {
      "extended": [
        "true",
        "false"
      ],
      "facing": [
        "north",
        "east",
        "south",
        "west",
        "up",
        "down"
      ]
    },
    "states": [
      {
        "properties": {
          "extended": "true",
          "facing": "north"
        },
        "id": 1028
      },
      {
        "properties": {
          "extended": "true",
          "facing": "east"
        },
        "id": 1029
      },
      {
        "properties": {
          "extended": "true",
          "facing": "south"
        },
        "id": 1030
      },
      {
        "properties": {
          "extended": "true",
          "facing": "west"
        },
        "id": 1031
      },
      {
        "properties": {
          "extended": "true",
          "facing": "up"
        },
        "id": 1032
      },
      {
        "properties": {
          "extended": "true",
          "facing": "down"
        },
        "id": 1033
      },
      {
        "properties": {
          "extended": "false",
          "facing": "north"
        },
        "id": 1034,
        "default": true
      },
      {
        "properties": {
          "extended": "false",
          "facing": "east"
        },
        "id": 1035
      },
      {
        "properties": {
          "extended": "false",
          "facing": "south"
        },
        "id": 1036
      },
      {
        "properties": {
          "extended": "false",
          "facing": "west"
        },
        "id": 1037
      },
      {
        "properties": {
          "extended": "false",
          "facing": "up"
        },
        "id": 1038
      },
      {
        "properties": {
          "extended": "false",
          "facing": "down"
        },
        "id": 1039
      }
    ]
  },
  "minecraft:cobweb": {
    "states": [
      {
        "id": 1040,
        "default": true
      }
    ]
  },
  "minecraft:grass": {
    "states": [
      {
        "id": 1041,
        "default": true
      }
    ]
  },
  "minecraft:fern": {
    "states": [
      {
        "id": 1042,
        "default": true
      }
    ]
  },
  "minecraft:dead_bush": {
    "states": [
      {
        "id": 1043,
        "default": true
      }
    ]
  },
  "minecraft:seagrass": {
    "states": [
      {
        "id": 1044,
        "default": true
      }
    ]
  },
  "minecraft:tall_seagrass": {
    "properties": {
      "half": [
        "upper",
        "lower"
      ]
    },
    "states": [
      {
        "properties": {
          "half": "upper"
        },
        "id": 1045
      },
      {
        "properties": {
          "half": "lower"
        },
        "id": 1046,
        "default": true
      }
    ]
  },
  "minecraft:piston": {
    "properties": {
      "extended": [
        "true",
        "false"
      ],
      "facing": [
        "north",
        "east",
        "south",
        "west",
        "up",
        "down"
      ]
    },
    "states": [
      {
        "properties": {
          "extended": "true",
          "facing": "north"
        },
        "id": 1047
      },
      {
        "properties": {
          "extended": "true",
          "facing": "east"
        },
        "id": 1048
      },
      {
        "properties": {
          "extended": "true",
          "facing": "south"
        },
        "id": 1049
      },
      {
        "properties": {
          "extended": "true",
          "facing": "west"
        },
        "id": 1050
      },
      {
        "properties": {
          "extended": "true",
          "facing": "up"
        },
        "id": 1051
      },
      {
        "properties": {
          "extended": "true",
          "facing": "down"
        },
        "id": 1052
      },
      {
        "properties": {
          "extended": "false",
          "facing": "north"
        },
        "id": 1053,
        "default": true
      },
      {
        "properties": {
          "extended": "false",
          "facing": "east"
        },
        "id": 1054
      },
      {
        "properties": {
          "extended": "false",
          "facing": "south"
        },
        "id": 1055
      },
      {
        "properties": {
          "extended": "false",
          "facing": "west"
        },
        "id": 1056
      },
      {
        "properties": {
          "extended": "false",
          "facing": "up"
        },
        "id": 1057
      },
      {
        "properties": {
          "extended": "false",
          "facing": "down"
        },
        "id": 1058
      }
    ]
  },
  "minecraft:piston_head": {
    "properties": {
      "facing": [
        "north",
        "east",
        "south",
        "west",
        "up",
        "down"
      ],
      "short": [
        "true",
        "false"
      ],
      "type": [
        "normal",
        "sticky"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "short": "true",
          "type": "normal"
        },
        "id": 1059
      },
      {
        "properties": {
          "facing": "north",
          "short": "true",
          "type": "sticky"
        },
        "id": 1060
      },
      {
        "properties": {
          "facing": "north",
          "short": "false",
          "type": "normal"
        },
        "id": 1061,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "short": "false",
          "type": "sticky"
        },
        "id": 1062
      },
      {
        "properties": {
          "facing": "east",
          "short": "true",
          "type": "normal"
        },
        "id": 1063
      },
      {
        "properties": {
          "facing": "east",
          "short": "true",
          "type": "sticky"
        },
        "id": 1064
      },
      {
        "properties": {
          "facing": "east",
          "short": "false",
          "type": "normal"
        },
        "id": 1065
      },
      {
        "properties": {
          "facing": "east",
          "short": "false",
          "type": "sticky"
        },
        "id": 1066
      },
      {
        "properties": {
          "facing": "south",
          "short": "true",
          "type": "normal"
        },
        "id": 1067
      },
      {
        "properties": {
          "facing": "south",
          "short": "true",
          "type": "sticky"
        },
        "id": 1068
      },
      {
        "properties": {
          "facing": "south",
          "short": "false",
          "type": "normal"
        },
        "id": 1069
      },
      {
        "properties": {
          "facing": "south",
          "short": "false",
          "type": "sticky"
        },
        "id": 1070
      },
      {
        "properties": {
          "facing": "west",
          "short": "true",
          "type": "normal"
        },
        "id": 1071
      },
      {
        "properties": {
          "facing": "west",
          "short": "true",
          "type": "sticky"
        },
        "id": 1072
      },
      {
        "properties": {
          "facing": "west",
          "short": "false",
          "type": "normal"
        },
        "id": 1073
      },
      {
        "properties": {
          "facing": "west",
          "short": "false",
          "type": "sticky"
        },
        "id": 1074
      },
      {
        "properties": {
          "facing": "up",
          "short": "true",
          "type": "normal"
        },
        "id": 1075
      },
      {
        "properties": {
          "facing": "up",
          "short": "true",
          "type": "sticky"
        },
        "id": 1076
      },
      {
        "properties": {
          "facing": "up",
          "short": "false",
          "type": "normal"
        },
        "id": 1077
      },
      {
        "properties": {
          "facing": "up",
          "short": "false",
          "type": "sticky"
        },
        "id": 1078
      },
      {
        "properties": {
          "facing": "down",
          "short": "true",
          "type": "normal"
        },
        "id": 1079
      },
      {
        "properties": {
          "facing": "down",
          "short": "true",
          "type": "sticky"
        },
        "id": 1080
      },
      {
        "properties": {
          "facing": "down",
          "short": "false",
          "type": "normal"
        },
        "id": 1081
      },
      {
        "properties": {
          "facing": "down",
          "short": "false",
          "type": "sticky"
        },
        "id": 1082
      }
    ]
  },
  "minecraft:white_wool": {
    "states": [
      {
        "id": 1083,
        "default": true
      }
    ]
  },
  "minecraft:orange_wool": {
    "states": [
      {
        "id": 1084,
        "default": true
      }
    ]
  },
  "minecraft:magenta_wool": {
    "states": [
      {
        "id": 1085,
        "default": true
      }
    ]
  },
  "minecraft:light_blue_wool": {
    "states": [
      {
        "id": 1086,
        "default": true
      }
    ]
  },
  "minecraft:yellow_wool": {
    "states": [
      {
        "id": 1087,
        "default": true
      }
    ]
  },
  "minecraft:lime_wool": {
    "states": [
      {
        "id": 1088,
        "default": true
      }
    ]
  },
  "minecraft:pink_wool": {
    "states": [
      {
        "id": 1089,
        "default": true
      }
    ]
  },
  "minecraft:gray_wool": {
    "states": [
      {
        "id": 1090,
        "default": true
      }
    ]
  },
  "minecraft:light_gray_wool": {
    "states": [
      {
        "id": 1091,
        "default": true
      }
    ]
  },
  "minecraft:cyan_wool": {
    "states": [
      {
        "id": 1092,
        "default": true
      }
    ]
  },
  "minecraft:purple_wool": {
    "states": [
      {
        "id": 1093,
        "default": true
      }
    ]
  },
  "minecraft:blue_wool": {
    "states": [
      {
        "id": 1094,
        "default": true
      }
    ]
  },
  "minecraft:brown_wool": {
    "states": [
      {
        "id": 1095,
        "default": true
      }
    ]
  },
  "minecraft:green_wool": {
    "states": [
      {
        "id": 1096,
        "default": true
      }
    ]
  },
  "minecraft:red_wool": {
    "states": [
      {
        "id": 1097,
        "default": true
      }
    ]
  },
  "minecraft:black_wool": {
    "states": [
      {
        "id": 1098,
        "default": true
      }
    ]
  },
  "minecraft:moving_piston": {
    "properties": {
      "facing": [
        "north",
        "east",
        "south",
        "west",
        "up",
        "down"
      ],
      "type": [
        "normal",
        "sticky"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "type": "normal"
        },
        "id": 1099,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "type": "sticky"
        },
        "id": 1100
      },
      {
        "properties": {
          "facing": "east",
          "type": "normal"
        },
        "id": 1101
      },
      {
        "properties": {
          "facing": "east",
          "type": "sticky"
        },
        "id": 1102
      },
      {
        "properties": {
          "facing": "south",
          "type": "normal"
        },
        "id": 1103
      },
      {
        "properties": {
          "facing": "south",
          "type": "sticky"
        },
        "id": 1104
      },
      {
        "properties": {
          "facing": "west",
          "type": "normal"
        },
        "id": 1105
      },
      {
        "properties": {
          "facing": "west",
          "type": "sticky"
        },
        "id": 1106
      },
      {
        "properties": {
          "facing": "up",
          "type": "normal"
        },
        "id": 1107
      },
      {
        "properties": {
          "facing": "up",
          "type": "sticky"
        },
        "id": 1108
      },
      {
        "properties": {
          "facing": "down",
          "type": "normal"
        },
        "id": 1109
      },
      {
        "properties": {
          "facing": "down",
          "type": "sticky"
        },
        "id": 1110
      }
    ]
  },
  "minecraft:dandelion": {
    "states": [
      {
        "id": 1111,
        "default": true
      }
    ]
  },
  "minecraft:poppy": {
    "states": [
      {
        "id": 1112,
        "default": true
      }
    ]
  },
  "minecraft:blue_orchid": {
    "states": [
      {
        "id": 1113,
        "default": true
      }
    ]
  },
  "minecraft:allium": {
    "states": [
      {
        "id": 1114,
        "default": true
      }
    ]
  },
  "minecraft:azure_bluet": {
    "states": [
      {
        "id": 1115,
        "default": true
      }
    ]
  },
  "minecraft:red_tulip": {
    "states": [
      {
        "id": 1116,
        "default": true
      }
    ]
  },
  "minecraft:orange_tulip": {
    "states": [
      {
        "id": 1117,
        "default": true
      }
    ]
  },
  "minecraft:white_tulip": {
    "states": [
      {
        "id": 1118,
        "default": true
      }
    ]
  },
  "minecraft:pink_tulip": {
    "states": [
      {
        "id": 1119,
        "default": true
      }
    ]
  },
  "minecraft:oxeye_daisy": {
    "states": [
      {
        "id": 1120,
        "default": true
      }
    ]
  },
  "minecraft:brown_mushroom": {
    "states": [
      {
        "id": 1121,
        "default": true
      }
    ]
  },
  "minecraft:red_mushroom": {
    "states": [
      {
        "id": 1122,
        "default": true
      }
    ]
  },
  "minecraft:gold_block": {
    "states": [
      {
        "id": 1123,
        "default": true
      }
    ]
  },
  "minecraft:iron_block": {
    "states": [
      {
        "id": 1124,
        "default": true
      }
    ]
  },
  "minecraft:bricks": {
    "states": [
      {
        "id": 1125,
        "default": true
      }
    ]
  },
  "minecraft:tnt": {
    "properties": {
      "unstable": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "unstable": "true"
        },
        "id": 1126
      },
      {
        "properties": {
          "unstable": "false"
        },
        "id": 1127,
        "default": true
      }
    ]
  },
  "minecraft:bookshelf": {
    "states": [
      {
        "id": 1128,
        "default": true
      }
    ]
  },
  "minecraft:mossy_cobblestone": {
    "states": [
      {
        "id": 1129,
        "default": true
      }
    ]
  },
  "minecraft:obsidian": {
    "states": [
      {
        "id": 1130,
        "default": true
      }
    ]
  },
  "minecraft:torch": {
    "states": [
      {
        "id": 1131,
        "default": true
      }
    ]
  },
  "minecraft:wall_torch": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north"
        },
        "id": 1132,
        "default": true
      },
      {
        "properties": {
          "facing": "south"
        },
        "id": 1133
      },
      {
        "properties": {
          "facing": "west"
        },
        "id": 1134
      },
      {
        "properties": {
          "facing": "east"
        },
        "id": 1135
      }
    ]
  },
  "minecraft:fire": {
    "properties": {
      "age": [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "10",
        "11",
        "12",
        "13",
        "14",
        "15"
      ],
      "east": [
        "true",
        "false"
      ],
      "north": [
        "true",
        "false"
      ],
      "south": [
        "true",
        "false"
      ],
      "up": [
        "true",
        "false"
      ],
      "west": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1136
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1137
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1138
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1139
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1140
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1141
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1142
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1143
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1144
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1145
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1146
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1147
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1148
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1149
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1150
      },
      {
        "properties": {
          "age": "0",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1151
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1152
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1153
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1154
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1155
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1156
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1157
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1158
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1159
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1160
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1161
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1162
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1163
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1164
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1165
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1166
      },
      {
        "properties": {
          "age": "0",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1167,
        "default": true
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1168
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1169
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1170
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1171
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1172
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1173
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1174
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1175
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1176
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1177
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1178
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1179
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1180
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1181
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1182
      },
      {
        "properties": {
          "age": "1",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1183
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1184
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1185
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1186
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1187
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1188
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1189
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1190
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1191
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1192
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1193
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1194
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1195
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1196
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1197
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1198
      },
      {
        "properties": {
          "age": "1",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1199
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1200
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1201
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1202
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1203
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1204
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1205
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1206
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1207
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1208
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1209
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1210
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1211
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1212
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1213
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1214
      },
      {
        "properties": {
          "age": "2",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1215
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1216
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1217
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1218
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1219
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1220
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1221
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1222
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1223
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1224
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1225
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1226
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1227
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1228
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1229
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1230
      },
      {
        "properties": {
          "age": "2",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1231
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1232
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1233
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1234
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1235
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1236
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1237
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1238
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1239
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1240
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1241
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1242
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1243
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1244
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1245
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1246
      },
      {
        "properties": {
          "age": "3",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1247
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1248
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1249
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1250
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1251
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1252
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1253
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1254
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1255
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1256
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1257
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1258
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1259
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1260
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1261
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1262
      },
      {
        "properties": {
          "age": "3",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1263
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1264
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1265
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1266
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1267
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1268
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1269
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1270
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1271
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1272
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1273
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1274
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1275
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1276
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1277
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1278
      },
      {
        "properties": {
          "age": "4",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1279
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1280
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1281
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1282
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1283
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1284
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1285
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1286
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1287
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1288
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1289
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1290
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1291
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1292
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1293
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1294
      },
      {
        "properties": {
          "age": "4",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1295
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1296
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1297
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1298
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1299
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1300
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1301
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1302
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1303
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1304
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1305
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1306
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1307
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1308
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1309
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1310
      },
      {
        "properties": {
          "age": "5",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1311
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1312
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1313
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1314
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1315
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1316
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1317
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1318
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1319
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1320
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1321
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1322
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1323
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1324
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1325
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1326
      },
      {
        "properties": {
          "age": "5",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1327
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1328
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1329
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1330
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1331
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1332
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1333
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1334
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1335
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1336
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1337
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1338
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1339
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1340
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1341
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1342
      },
      {
        "properties": {
          "age": "6",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1343
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1344
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1345
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1346
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1347
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1348
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1349
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1350
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1351
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1352
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1353
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1354
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1355
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1356
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1357
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1358
      },
      {
        "properties": {
          "age": "6",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1359
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1360
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1361
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1362
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1363
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1364
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1365
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1366
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1367
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1368
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1369
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1370
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1371
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1372
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1373
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1374
      },
      {
        "properties": {
          "age": "7",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1375
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1376
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1377
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1378
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1379
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1380
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1381
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1382
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1383
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1384
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1385
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1386
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1387
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1388
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1389
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1390
      },
      {
        "properties": {
          "age": "7",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1391
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1392
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1393
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1394
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1395
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1396
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1397
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1398
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1399
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1400
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1401
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1402
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1403
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1404
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1405
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1406
      },
      {
        "properties": {
          "age": "8",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1407
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1408
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1409
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1410
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1411
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1412
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1413
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1414
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1415
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1416
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1417
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1418
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1419
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1420
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1421
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1422
      },
      {
        "properties": {
          "age": "8",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1423
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1424
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1425
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1426
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1427
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1428
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1429
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1430
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1431
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1432
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1433
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1434
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1435
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1436
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1437
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1438
      },
      {
        "properties": {
          "age": "9",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1439
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1440
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1441
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1442
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1443
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1444
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1445
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1446
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1447
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1448
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1449
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1450
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1451
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1452
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1453
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1454
      },
      {
        "properties": {
          "age": "9",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1455
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1456
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1457
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1458
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1459
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1460
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1461
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1462
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1463
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1464
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1465
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1466
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1467
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1468
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1469
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1470
      },
      {
        "properties": {
          "age": "10",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1471
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1472
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1473
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1474
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1475
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1476
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1477
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1478
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1479
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1480
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1481
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1482
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1483
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1484
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1485
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1486
      },
      {
        "properties": {
          "age": "10",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1487
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1488
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1489
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1490
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1491
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1492
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1493
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1494
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1495
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1496
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1497
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1498
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1499
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1500
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1501
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1502
      },
      {
        "properties": {
          "age": "11",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1503
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1504
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1505
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1506
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1507
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1508
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1509
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1510
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1511
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1512
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1513
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1514
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1515
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1516
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1517
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1518
      },
      {
        "properties": {
          "age": "11",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1519
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1520
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1521
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1522
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1523
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1524
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1525
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1526
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1527
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1528
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1529
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1530
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1531
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1532
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1533
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1534
      },
      {
        "properties": {
          "age": "12",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1535
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1536
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1537
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1538
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1539
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1540
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1541
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1542
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1543
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1544
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1545
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1546
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1547
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1548
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1549
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1550
      },
      {
        "properties": {
          "age": "12",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1551
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1552
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1553
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1554
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1555
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1556
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1557
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1558
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1559
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1560
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1561
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1562
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1563
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1564
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1565
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1566
      },
      {
        "properties": {
          "age": "13",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1567
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1568
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1569
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1570
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1571
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1572
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1573
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1574
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1575
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1576
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1577
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1578
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1579
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1580
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1581
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1582
      },
      {
        "properties": {
          "age": "13",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1583
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1584
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1585
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1586
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1587
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1588
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1589
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1590
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1591
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1592
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1593
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1594
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1595
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1596
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1597
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1598
      },
      {
        "properties": {
          "age": "14",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1599
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1600
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1601
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1602
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1603
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1604
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1605
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1606
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1607
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1608
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1609
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1610
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1611
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1612
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1613
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1614
      },
      {
        "properties": {
          "age": "14",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1615
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1616
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1617
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1618
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1619
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1620
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1621
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1622
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1623
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1624
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1625
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1626
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1627
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1628
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1629
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1630
      },
      {
        "properties": {
          "age": "15",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1631
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1632
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1633
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1634
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1635
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1636
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1637
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1638
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1639
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 1640
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 1641
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 1642
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 1643
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 1644
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 1645
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 1646
      },
      {
        "properties": {
          "age": "15",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 1647
      }
    ]
  },
  "minecraft:spawner": {
    "states": [
      {
        "id": 1648,
        "default": true
      }
    ]
  },
  "minecraft:oak_stairs": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "shape": [
        "straight",
        "inner_left",
        "inner_right",
        "outer_left",
        "outer_right"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 1649
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 1650
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 1651
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 1652
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 1653
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 1654
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 1655
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 1656
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 1657
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 1658
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 1659
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 1660,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 1661
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 1662
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 1663
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 1664
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 1665
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 1666
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 1667
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 1668
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 1669
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 1670
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 1671
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 1672
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 1673
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 1674
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 1675
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 1676
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 1677
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 1678
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 1679
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 1680
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 1681
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 1682
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 1683
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 1684
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 1685
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 1686
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 1687
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 1688
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 1689
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 1690
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 1691
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 1692
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 1693
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 1694
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 1695
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 1696
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 1697
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 1698
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 1699
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 1700
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 1701
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 1702
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 1703
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 1704
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 1705
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 1706
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 1707
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 1708
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 1709
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 1710
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 1711
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 1712
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 1713
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 1714
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 1715
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 1716
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 1717
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 1718
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 1719
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 1720
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 1721
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 1722
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 1723
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 1724
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 1725
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 1726
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 1727
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 1728
      }
    ]
  },
  "minecraft:chest": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "type": [
        "single",
        "left",
        "right"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "type": "single",
          "waterlogged": "true"
        },
        "id": 1729
      },
      {
        "properties": {
          "facing": "north",
          "type": "single",
          "waterlogged": "false"
        },
        "id": 1730,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "type": "left",
          "waterlogged": "true"
        },
        "id": 1731
      },
      {
        "properties": {
          "facing": "north",
          "type": "left",
          "waterlogged": "false"
        },
        "id": 1732
      },
      {
        "properties": {
          "facing": "north",
          "type": "right",
          "waterlogged": "true"
        },
        "id": 1733
      },
      {
        "properties": {
          "facing": "north",
          "type": "right",
          "waterlogged": "false"
        },
        "id": 1734
      },
      {
        "properties": {
          "facing": "south",
          "type": "single",
          "waterlogged": "true"
        },
        "id": 1735
      },
      {
        "properties": {
          "facing": "south",
          "type": "single",
          "waterlogged": "false"
        },
        "id": 1736
      },
      {
        "properties": {
          "facing": "south",
          "type": "left",
          "waterlogged": "true"
        },
        "id": 1737
      },
      {
        "properties": {
          "facing": "south",
          "type": "left",
          "waterlogged": "false"
        },
        "id": 1738
      },
      {
        "properties": {
          "facing": "south",
          "type": "right",
          "waterlogged": "true"
        },
        "id": 1739
      },
      {
        "properties": {
          "facing": "south",
          "type": "right",
          "waterlogged": "false"
        },
        "id": 1740
      },
      {
        "properties": {
          "facing": "west",
          "type": "single",
          "waterlogged": "true"
        },
        "id": 1741
      },
      {
        "properties": {
          "facing": "west",
          "type": "single",
          "waterlogged": "false"
        },
        "id": 1742
      },
      {
        "properties": {
          "facing": "west",
          "type": "left",
          "waterlogged": "true"
        },
        "id": 1743
      },
      {
        "properties": {
          "facing": "west",
          "type": "left",
          "waterlogged": "false"
        },
        "id": 1744
      },
      {
        "properties": {
          "facing": "west",
          "type": "right",
          "waterlogged": "true"
        },
        "id": 1745
      },
      {
        "properties": {
          "facing": "west",
          "type": "right",
          "waterlogged": "false"
        },
        "id": 1746
      },
      {
        "properties": {
          "facing": "east",
          "type": "single",
          "waterlogged": "true"
        },
        "id": 1747
      },
      {
        "properties": {
          "facing": "east",
          "type": "single",
          "waterlogged": "false"
        },
        "id": 1748
      },
      {
        "properties": {
          "facing": "east",
          "type": "left",
          "waterlogged": "true"
        },
        "id": 1749
      },
      {
        "properties": {
          "facing": "east",
          "type": "left",
          "waterlogged": "false"
        },
        "id": 1750
      },
      {
        "properties": {
          "facing": "east",
          "type": "right",
          "waterlogged": "true"
        },
        "id": 1751
      },
      {
        "properties": {
          "facing": "east",
          "type": "right",
          "waterlogged": "false"
        },
        "id": 1752
      }
    ]
  },
  "minecraft:redstone_wire": {
    "properties": {
      "east": [
        "up",
        "side",
        "none"
      ],
      "north": [
        "up",
        "side",
        "none"
      ],
      "power": [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "10",
        "11",
        "12",
        "13",
        "14",
        "15"
      ],
      "south": [
        "up",
        "side",
        "none"
      ],
      "west": [
        "up",
        "side",
        "none"
      ]
    },
    "states": [
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "0",
          "south": "up",
          "west": "up"
        },
        "id": 1753
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "0",
          "south": "up",
          "west": "side"
        },
        "id": 1754
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "0",
          "south": "up",
          "west": "none"
        },
        "id": 1755
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "0",
          "south": "side",
          "west": "up"
        },
        "id": 1756
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "0",
          "south": "side",
          "west": "side"
        },
        "id": 1757
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "0",
          "south": "side",
          "west": "none"
        },
        "id": 1758
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "0",
          "south": "none",
          "west": "up"
        },
        "id": 1759
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "0",
          "south": "none",
          "west": "side"
        },
        "id": 1760
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "0",
          "south": "none",
          "west": "none"
        },
        "id": 1761
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "1",
          "south": "up",
          "west": "up"
        },
        "id": 1762
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "1",
          "south": "up",
          "west": "side"
        },
        "id": 1763
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "1",
          "south": "up",
          "west": "none"
        },
        "id": 1764
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "1",
          "south": "side",
          "west": "up"
        },
        "id": 1765
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "1",
          "south": "side",
          "west": "side"
        },
        "id": 1766
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "1",
          "south": "side",
          "west": "none"
        },
        "id": 1767
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "1",
          "south": "none",
          "west": "up"
        },
        "id": 1768
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "1",
          "south": "none",
          "west": "side"
        },
        "id": 1769
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "1",
          "south": "none",
          "west": "none"
        },
        "id": 1770
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "2",
          "south": "up",
          "west": "up"
        },
        "id": 1771
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "2",
          "south": "up",
          "west": "side"
        },
        "id": 1772
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "2",
          "south": "up",
          "west": "none"
        },
        "id": 1773
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "2",
          "south": "side",
          "west": "up"
        },
        "id": 1774
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "2",
          "south": "side",
          "west": "side"
        },
        "id": 1775
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "2",
          "south": "side",
          "west": "none"
        },
        "id": 1776
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "2",
          "south": "none",
          "west": "up"
        },
        "id": 1777
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "2",
          "south": "none",
          "west": "side"
        },
        "id": 1778
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "2",
          "south": "none",
          "west": "none"
        },
        "id": 1779
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "3",
          "south": "up",
          "west": "up"
        },
        "id": 1780
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "3",
          "south": "up",
          "west": "side"
        },
        "id": 1781
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "3",
          "south": "up",
          "west": "none"
        },
        "id": 1782
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "3",
          "south": "side",
          "west": "up"
        },
        "id": 1783
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "3",
          "south": "side",
          "west": "side"
        },
        "id": 1784
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "3",
          "south": "side",
          "west": "none"
        },
        "id": 1785
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "3",
          "south": "none",
          "west": "up"
        },
        "id": 1786
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "3",
          "south": "none",
          "west": "side"
        },
        "id": 1787
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "3",
          "south": "none",
          "west": "none"
        },
        "id": 1788
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "4",
          "south": "up",
          "west": "up"
        },
        "id": 1789
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "4",
          "south": "up",
          "west": "side"
        },
        "id": 1790
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "4",
          "south": "up",
          "west": "none"
        },
        "id": 1791
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "4",
          "south": "side",
          "west": "up"
        },
        "id": 1792
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "4",
          "south": "side",
          "west": "side"
        },
        "id": 1793
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "4",
          "south": "side",
          "west": "none"
        },
        "id": 1794
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "4",
          "south": "none",
          "west": "up"
        },
        "id": 1795
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "4",
          "south": "none",
          "west": "side"
        },
        "id": 1796
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "4",
          "south": "none",
          "west": "none"
        },
        "id": 1797
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "5",
          "south": "up",
          "west": "up"
        },
        "id": 1798
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "5",
          "south": "up",
          "west": "side"
        },
        "id": 1799
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "5",
          "south": "up",
          "west": "none"
        },
        "id": 1800
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "5",
          "south": "side",
          "west": "up"
        },
        "id": 1801
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "5",
          "south": "side",
          "west": "side"
        },
        "id": 1802
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "5",
          "south": "side",
          "west": "none"
        },
        "id": 1803
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "5",
          "south": "none",
          "west": "up"
        },
        "id": 1804
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "5",
          "south": "none",
          "west": "side"
        },
        "id": 1805
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "5",
          "south": "none",
          "west": "none"
        },
        "id": 1806
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "6",
          "south": "up",
          "west": "up"
        },
        "id": 1807
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "6",
          "south": "up",
          "west": "side"
        },
        "id": 1808
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "6",
          "south": "up",
          "west": "none"
        },
        "id": 1809
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "6",
          "south": "side",
          "west": "up"
        },
        "id": 1810
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "6",
          "south": "side",
          "west": "side"
        },
        "id": 1811
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "6",
          "south": "side",
          "west": "none"
        },
        "id": 1812
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "6",
          "south": "none",
          "west": "up"
        },
        "id": 1813
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "6",
          "south": "none",
          "west": "side"
        },
        "id": 1814
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "6",
          "south": "none",
          "west": "none"
        },
        "id": 1815
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "7",
          "south": "up",
          "west": "up"
        },
        "id": 1816
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "7",
          "south": "up",
          "west": "side"
        },
        "id": 1817
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "7",
          "south": "up",
          "west": "none"
        },
        "id": 1818
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "7",
          "south": "side",
          "west": "up"
        },
        "id": 1819
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "7",
          "south": "side",
          "west": "side"
        },
        "id": 1820
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "7",
          "south": "side",
          "west": "none"
        },
        "id": 1821
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "7",
          "south": "none",
          "west": "up"
        },
        "id": 1822
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "7",
          "south": "none",
          "west": "side"
        },
        "id": 1823
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "7",
          "south": "none",
          "west": "none"
        },
        "id": 1824
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "8",
          "south": "up",
          "west": "up"
        },
        "id": 1825
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "8",
          "south": "up",
          "west": "side"
        },
        "id": 1826
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "8",
          "south": "up",
          "west": "none"
        },
        "id": 1827
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "8",
          "south": "side",
          "west": "up"
        },
        "id": 1828
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "8",
          "south": "side",
          "west": "side"
        },
        "id": 1829
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "8",
          "south": "side",
          "west": "none"
        },
        "id": 1830
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "8",
          "south": "none",
          "west": "up"
        },
        "id": 1831
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "8",
          "south": "none",
          "west": "side"
        },
        "id": 1832
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "8",
          "south": "none",
          "west": "none"
        },
        "id": 1833
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "9",
          "south": "up",
          "west": "up"
        },
        "id": 1834
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "9",
          "south": "up",
          "west": "side"
        },
        "id": 1835
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "9",
          "south": "up",
          "west": "none"
        },
        "id": 1836
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "9",
          "south": "side",
          "west": "up"
        },
        "id": 1837
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "9",
          "south": "side",
          "west": "side"
        },
        "id": 1838
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "9",
          "south": "side",
          "west": "none"
        },
        "id": 1839
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "9",
          "south": "none",
          "west": "up"
        },
        "id": 1840
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "9",
          "south": "none",
          "west": "side"
        },
        "id": 1841
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "9",
          "south": "none",
          "west": "none"
        },
        "id": 1842
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "10",
          "south": "up",
          "west": "up"
        },
        "id": 1843
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "10",
          "south": "up",
          "west": "side"
        },
        "id": 1844
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "10",
          "south": "up",
          "west": "none"
        },
        "id": 1845
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "10",
          "south": "side",
          "west": "up"
        },
        "id": 1846
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "10",
          "south": "side",
          "west": "side"
        },
        "id": 1847
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "10",
          "south": "side",
          "west": "none"
        },
        "id": 1848
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "10",
          "south": "none",
          "west": "up"
        },
        "id": 1849
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "10",
          "south": "none",
          "west": "side"
        },
        "id": 1850
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "10",
          "south": "none",
          "west": "none"
        },
        "id": 1851
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "11",
          "south": "up",
          "west": "up"
        },
        "id": 1852
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "11",
          "south": "up",
          "west": "side"
        },
        "id": 1853
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "11",
          "south": "up",
          "west": "none"
        },
        "id": 1854
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "11",
          "south": "side",
          "west": "up"
        },
        "id": 1855
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "11",
          "south": "side",
          "west": "side"
        },
        "id": 1856
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "11",
          "south": "side",
          "west": "none"
        },
        "id": 1857
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "11",
          "south": "none",
          "west": "up"
        },
        "id": 1858
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "11",
          "south": "none",
          "west": "side"
        },
        "id": 1859
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "11",
          "south": "none",
          "west": "none"
        },
        "id": 1860
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "12",
          "south": "up",
          "west": "up"
        },
        "id": 1861
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "12",
          "south": "up",
          "west": "side"
        },
        "id": 1862
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "12",
          "south": "up",
          "west": "none"
        },
        "id": 1863
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "12",
          "south": "side",
          "west": "up"
        },
        "id": 1864
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "12",
          "south": "side",
          "west": "side"
        },
        "id": 1865
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "12",
          "south": "side",
          "west": "none"
        },
        "id": 1866
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "12",
          "south": "none",
          "west": "up"
        },
        "id": 1867
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "12",
          "south": "none",
          "west": "side"
        },
        "id": 1868
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "12",
          "south": "none",
          "west": "none"
        },
        "id": 1869
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "13",
          "south": "up",
          "west": "up"
        },
        "id": 1870
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "13",
          "south": "up",
          "west": "side"
        },
        "id": 1871
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "13",
          "south": "up",
          "west": "none"
        },
        "id": 1872
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "13",
          "south": "side",
          "west": "up"
        },
        "id": 1873
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "13",
          "south": "side",
          "west": "side"
        },
        "id": 1874
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "13",
          "south": "side",
          "west": "none"
        },
        "id": 1875
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "13",
          "south": "none",
          "west": "up"
        },
        "id": 1876
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "13",
          "south": "none",
          "west": "side"
        },
        "id": 1877
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "13",
          "south": "none",
          "west": "none"
        },
        "id": 1878
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "14",
          "south": "up",
          "west": "up"
        },
        "id": 1879
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "14",
          "south": "up",
          "west": "side"
        },
        "id": 1880
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "14",
          "south": "up",
          "west": "none"
        },
        "id": 1881
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "14",
          "south": "side",
          "west": "up"
        },
        "id": 1882
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "14",
          "south": "side",
          "west": "side"
        },
        "id": 1883
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "14",
          "south": "side",
          "west": "none"
        },
        "id": 1884
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "14",
          "south": "none",
          "west": "up"
        },
        "id": 1885
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "14",
          "south": "none",
          "west": "side"
        },
        "id": 1886
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "14",
          "south": "none",
          "west": "none"
        },
        "id": 1887
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "15",
          "south": "up",
          "west": "up"
        },
        "id": 1888
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "15",
          "south": "up",
          "west": "side"
        },
        "id": 1889
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "15",
          "south": "up",
          "west": "none"
        },
        "id": 1890
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "15",
          "south": "side",
          "west": "up"
        },
        "id": 1891
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "15",
          "south": "side",
          "west": "side"
        },
        "id": 1892
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "15",
          "south": "side",
          "west": "none"
        },
        "id": 1893
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "15",
          "south": "none",
          "west": "up"
        },
        "id": 1894
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "15",
          "south": "none",
          "west": "side"
        },
        "id": 1895
      },
      {
        "properties": {
          "east": "up",
          "north": "up",
          "power": "15",
          "south": "none",
          "west": "none"
        },
        "id": 1896
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "0",
          "south": "up",
          "west": "up"
        },
        "id": 1897
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "0",
          "south": "up",
          "west": "side"
        },
        "id": 1898
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "0",
          "south": "up",
          "west": "none"
        },
        "id": 1899
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "0",
          "south": "side",
          "west": "up"
        },
        "id": 1900
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "0",
          "south": "side",
          "west": "side"
        },
        "id": 1901
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "0",
          "south": "side",
          "west": "none"
        },
        "id": 1902
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "0",
          "south": "none",
          "west": "up"
        },
        "id": 1903
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "0",
          "south": "none",
          "west": "side"
        },
        "id": 1904
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "0",
          "south": "none",
          "west": "none"
        },
        "id": 1905
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "1",
          "south": "up",
          "west": "up"
        },
        "id": 1906
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "1",
          "south": "up",
          "west": "side"
        },
        "id": 1907
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "1",
          "south": "up",
          "west": "none"
        },
        "id": 1908
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "1",
          "south": "side",
          "west": "up"
        },
        "id": 1909
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "1",
          "south": "side",
          "west": "side"
        },
        "id": 1910
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "1",
          "south": "side",
          "west": "none"
        },
        "id": 1911
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "1",
          "south": "none",
          "west": "up"
        },
        "id": 1912
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "1",
          "south": "none",
          "west": "side"
        },
        "id": 1913
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "1",
          "south": "none",
          "west": "none"
        },
        "id": 1914
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "2",
          "south": "up",
          "west": "up"
        },
        "id": 1915
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "2",
          "south": "up",
          "west": "side"
        },
        "id": 1916
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "2",
          "south": "up",
          "west": "none"
        },
        "id": 1917
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "2",
          "south": "side",
          "west": "up"
        },
        "id": 1918
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "2",
          "south": "side",
          "west": "side"
        },
        "id": 1919
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "2",
          "south": "side",
          "west": "none"
        },
        "id": 1920
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "2",
          "south": "none",
          "west": "up"
        },
        "id": 1921
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "2",
          "south": "none",
          "west": "side"
        },
        "id": 1922
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "2",
          "south": "none",
          "west": "none"
        },
        "id": 1923
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "3",
          "south": "up",
          "west": "up"
        },
        "id": 1924
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "3",
          "south": "up",
          "west": "side"
        },
        "id": 1925
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "3",
          "south": "up",
          "west": "none"
        },
        "id": 1926
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "3",
          "south": "side",
          "west": "up"
        },
        "id": 1927
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "3",
          "south": "side",
          "west": "side"
        },
        "id": 1928
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "3",
          "south": "side",
          "west": "none"
        },
        "id": 1929
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "3",
          "south": "none",
          "west": "up"
        },
        "id": 1930
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "3",
          "south": "none",
          "west": "side"
        },
        "id": 1931
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "3",
          "south": "none",
          "west": "none"
        },
        "id": 1932
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "4",
          "south": "up",
          "west": "up"
        },
        "id": 1933
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "4",
          "south": "up",
          "west": "side"
        },
        "id": 1934
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "4",
          "south": "up",
          "west": "none"
        },
        "id": 1935
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "4",
          "south": "side",
          "west": "up"
        },
        "id": 1936
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "4",
          "south": "side",
          "west": "side"
        },
        "id": 1937
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "4",
          "south": "side",
          "west": "none"
        },
        "id": 1938
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "4",
          "south": "none",
          "west": "up"
        },
        "id": 1939
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "4",
          "south": "none",
          "west": "side"
        },
        "id": 1940
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "4",
          "south": "none",
          "west": "none"
        },
        "id": 1941
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "5",
          "south": "up",
          "west": "up"
        },
        "id": 1942
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "5",
          "south": "up",
          "west": "side"
        },
        "id": 1943
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "5",
          "south": "up",
          "west": "none"
        },
        "id": 1944
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "5",
          "south": "side",
          "west": "up"
        },
        "id": 1945
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "5",
          "south": "side",
          "west": "side"
        },
        "id": 1946
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "5",
          "south": "side",
          "west": "none"
        },
        "id": 1947
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "5",
          "south": "none",
          "west": "up"
        },
        "id": 1948
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "5",
          "south": "none",
          "west": "side"
        },
        "id": 1949
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "5",
          "south": "none",
          "west": "none"
        },
        "id": 1950
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "6",
          "south": "up",
          "west": "up"
        },
        "id": 1951
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "6",
          "south": "up",
          "west": "side"
        },
        "id": 1952
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "6",
          "south": "up",
          "west": "none"
        },
        "id": 1953
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "6",
          "south": "side",
          "west": "up"
        },
        "id": 1954
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "6",
          "south": "side",
          "west": "side"
        },
        "id": 1955
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "6",
          "south": "side",
          "west": "none"
        },
        "id": 1956
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "6",
          "south": "none",
          "west": "up"
        },
        "id": 1957
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "6",
          "south": "none",
          "west": "side"
        },
        "id": 1958
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "6",
          "south": "none",
          "west": "none"
        },
        "id": 1959
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "7",
          "south": "up",
          "west": "up"
        },
        "id": 1960
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "7",
          "south": "up",
          "west": "side"
        },
        "id": 1961
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "7",
          "south": "up",
          "west": "none"
        },
        "id": 1962
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "7",
          "south": "side",
          "west": "up"
        },
        "id": 1963
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "7",
          "south": "side",
          "west": "side"
        },
        "id": 1964
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "7",
          "south": "side",
          "west": "none"
        },
        "id": 1965
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "7",
          "south": "none",
          "west": "up"
        },
        "id": 1966
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "7",
          "south": "none",
          "west": "side"
        },
        "id": 1967
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "7",
          "south": "none",
          "west": "none"
        },
        "id": 1968
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "8",
          "south": "up",
          "west": "up"
        },
        "id": 1969
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "8",
          "south": "up",
          "west": "side"
        },
        "id": 1970
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "8",
          "south": "up",
          "west": "none"
        },
        "id": 1971
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "8",
          "south": "side",
          "west": "up"
        },
        "id": 1972
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "8",
          "south": "side",
          "west": "side"
        },
        "id": 1973
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "8",
          "south": "side",
          "west": "none"
        },
        "id": 1974
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "8",
          "south": "none",
          "west": "up"
        },
        "id": 1975
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "8",
          "south": "none",
          "west": "side"
        },
        "id": 1976
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "8",
          "south": "none",
          "west": "none"
        },
        "id": 1977
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "9",
          "south": "up",
          "west": "up"
        },
        "id": 1978
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "9",
          "south": "up",
          "west": "side"
        },
        "id": 1979
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "9",
          "south": "up",
          "west": "none"
        },
        "id": 1980
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "9",
          "south": "side",
          "west": "up"
        },
        "id": 1981
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "9",
          "south": "side",
          "west": "side"
        },
        "id": 1982
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "9",
          "south": "side",
          "west": "none"
        },
        "id": 1983
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "9",
          "south": "none",
          "west": "up"
        },
        "id": 1984
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "9",
          "south": "none",
          "west": "side"
        },
        "id": 1985
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "9",
          "south": "none",
          "west": "none"
        },
        "id": 1986
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "10",
          "south": "up",
          "west": "up"
        },
        "id": 1987
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "10",
          "south": "up",
          "west": "side"
        },
        "id": 1988
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "10",
          "south": "up",
          "west": "none"
        },
        "id": 1989
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "10",
          "south": "side",
          "west": "up"
        },
        "id": 1990
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "10",
          "south": "side",
          "west": "side"
        },
        "id": 1991
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "10",
          "south": "side",
          "west": "none"
        },
        "id": 1992
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "10",
          "south": "none",
          "west": "up"
        },
        "id": 1993
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "10",
          "south": "none",
          "west": "side"
        },
        "id": 1994
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "10",
          "south": "none",
          "west": "none"
        },
        "id": 1995
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "11",
          "south": "up",
          "west": "up"
        },
        "id": 1996
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "11",
          "south": "up",
          "west": "side"
        },
        "id": 1997
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "11",
          "south": "up",
          "west": "none"
        },
        "id": 1998
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "11",
          "south": "side",
          "west": "up"
        },
        "id": 1999
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "11",
          "south": "side",
          "west": "side"
        },
        "id": 2000
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "11",
          "south": "side",
          "west": "none"
        },
        "id": 2001
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "11",
          "south": "none",
          "west": "up"
        },
        "id": 2002
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "11",
          "south": "none",
          "west": "side"
        },
        "id": 2003
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "11",
          "south": "none",
          "west": "none"
        },
        "id": 2004
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "12",
          "south": "up",
          "west": "up"
        },
        "id": 2005
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "12",
          "south": "up",
          "west": "side"
        },
        "id": 2006
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "12",
          "south": "up",
          "west": "none"
        },
        "id": 2007
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "12",
          "south": "side",
          "west": "up"
        },
        "id": 2008
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "12",
          "south": "side",
          "west": "side"
        },
        "id": 2009
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "12",
          "south": "side",
          "west": "none"
        },
        "id": 2010
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "12",
          "south": "none",
          "west": "up"
        },
        "id": 2011
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "12",
          "south": "none",
          "west": "side"
        },
        "id": 2012
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "12",
          "south": "none",
          "west": "none"
        },
        "id": 2013
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "13",
          "south": "up",
          "west": "up"
        },
        "id": 2014
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "13",
          "south": "up",
          "west": "side"
        },
        "id": 2015
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "13",
          "south": "up",
          "west": "none"
        },
        "id": 2016
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "13",
          "south": "side",
          "west": "up"
        },
        "id": 2017
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "13",
          "south": "side",
          "west": "side"
        },
        "id": 2018
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "13",
          "south": "side",
          "west": "none"
        },
        "id": 2019
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "13",
          "south": "none",
          "west": "up"
        },
        "id": 2020
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "13",
          "south": "none",
          "west": "side"
        },
        "id": 2021
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "13",
          "south": "none",
          "west": "none"
        },
        "id": 2022
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "14",
          "south": "up",
          "west": "up"
        },
        "id": 2023
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "14",
          "south": "up",
          "west": "side"
        },
        "id": 2024
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "14",
          "south": "up",
          "west": "none"
        },
        "id": 2025
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "14",
          "south": "side",
          "west": "up"
        },
        "id": 2026
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "14",
          "south": "side",
          "west": "side"
        },
        "id": 2027
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "14",
          "south": "side",
          "west": "none"
        },
        "id": 2028
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "14",
          "south": "none",
          "west": "up"
        },
        "id": 2029
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "14",
          "south": "none",
          "west": "side"
        },
        "id": 2030
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "14",
          "south": "none",
          "west": "none"
        },
        "id": 2031
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "15",
          "south": "up",
          "west": "up"
        },
        "id": 2032
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "15",
          "south": "up",
          "west": "side"
        },
        "id": 2033
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "15",
          "south": "up",
          "west": "none"
        },
        "id": 2034
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "15",
          "south": "side",
          "west": "up"
        },
        "id": 2035
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "15",
          "south": "side",
          "west": "side"
        },
        "id": 2036
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "15",
          "south": "side",
          "west": "none"
        },
        "id": 2037
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "15",
          "south": "none",
          "west": "up"
        },
        "id": 2038
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "15",
          "south": "none",
          "west": "side"
        },
        "id": 2039
      },
      {
        "properties": {
          "east": "up",
          "north": "side",
          "power": "15",
          "south": "none",
          "west": "none"
        },
        "id": 2040
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "0",
          "south": "up",
          "west": "up"
        },
        "id": 2041
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "0",
          "south": "up",
          "west": "side"
        },
        "id": 2042
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "0",
          "south": "up",
          "west": "none"
        },
        "id": 2043
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "0",
          "south": "side",
          "west": "up"
        },
        "id": 2044
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "0",
          "south": "side",
          "west": "side"
        },
        "id": 2045
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "0",
          "south": "side",
          "west": "none"
        },
        "id": 2046
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "0",
          "south": "none",
          "west": "up"
        },
        "id": 2047
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "0",
          "south": "none",
          "west": "side"
        },
        "id": 2048
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "0",
          "south": "none",
          "west": "none"
        },
        "id": 2049
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "1",
          "south": "up",
          "west": "up"
        },
        "id": 2050
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "1",
          "south": "up",
          "west": "side"
        },
        "id": 2051
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "1",
          "south": "up",
          "west": "none"
        },
        "id": 2052
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "1",
          "south": "side",
          "west": "up"
        },
        "id": 2053
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "1",
          "south": "side",
          "west": "side"
        },
        "id": 2054
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "1",
          "south": "side",
          "west": "none"
        },
        "id": 2055
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "1",
          "south": "none",
          "west": "up"
        },
        "id": 2056
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "1",
          "south": "none",
          "west": "side"
        },
        "id": 2057
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "1",
          "south": "none",
          "west": "none"
        },
        "id": 2058
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "2",
          "south": "up",
          "west": "up"
        },
        "id": 2059
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "2",
          "south": "up",
          "west": "side"
        },
        "id": 2060
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "2",
          "south": "up",
          "west": "none"
        },
        "id": 2061
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "2",
          "south": "side",
          "west": "up"
        },
        "id": 2062
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "2",
          "south": "side",
          "west": "side"
        },
        "id": 2063
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "2",
          "south": "side",
          "west": "none"
        },
        "id": 2064
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "2",
          "south": "none",
          "west": "up"
        },
        "id": 2065
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "2",
          "south": "none",
          "west": "side"
        },
        "id": 2066
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "2",
          "south": "none",
          "west": "none"
        },
        "id": 2067
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "3",
          "south": "up",
          "west": "up"
        },
        "id": 2068
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "3",
          "south": "up",
          "west": "side"
        },
        "id": 2069
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "3",
          "south": "up",
          "west": "none"
        },
        "id": 2070
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "3",
          "south": "side",
          "west": "up"
        },
        "id": 2071
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "3",
          "south": "side",
          "west": "side"
        },
        "id": 2072
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "3",
          "south": "side",
          "west": "none"
        },
        "id": 2073
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "3",
          "south": "none",
          "west": "up"
        },
        "id": 2074
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "3",
          "south": "none",
          "west": "side"
        },
        "id": 2075
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "3",
          "south": "none",
          "west": "none"
        },
        "id": 2076
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "4",
          "south": "up",
          "west": "up"
        },
        "id": 2077
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "4",
          "south": "up",
          "west": "side"
        },
        "id": 2078
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "4",
          "south": "up",
          "west": "none"
        },
        "id": 2079
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "4",
          "south": "side",
          "west": "up"
        },
        "id": 2080
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "4",
          "south": "side",
          "west": "side"
        },
        "id": 2081
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "4",
          "south": "side",
          "west": "none"
        },
        "id": 2082
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "4",
          "south": "none",
          "west": "up"
        },
        "id": 2083
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "4",
          "south": "none",
          "west": "side"
        },
        "id": 2084
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "4",
          "south": "none",
          "west": "none"
        },
        "id": 2085
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "5",
          "south": "up",
          "west": "up"
        },
        "id": 2086
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "5",
          "south": "up",
          "west": "side"
        },
        "id": 2087
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "5",
          "south": "up",
          "west": "none"
        },
        "id": 2088
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "5",
          "south": "side",
          "west": "up"
        },
        "id": 2089
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "5",
          "south": "side",
          "west": "side"
        },
        "id": 2090
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "5",
          "south": "side",
          "west": "none"
        },
        "id": 2091
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "5",
          "south": "none",
          "west": "up"
        },
        "id": 2092
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "5",
          "south": "none",
          "west": "side"
        },
        "id": 2093
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "5",
          "south": "none",
          "west": "none"
        },
        "id": 2094
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "6",
          "south": "up",
          "west": "up"
        },
        "id": 2095
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "6",
          "south": "up",
          "west": "side"
        },
        "id": 2096
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "6",
          "south": "up",
          "west": "none"
        },
        "id": 2097
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "6",
          "south": "side",
          "west": "up"
        },
        "id": 2098
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "6",
          "south": "side",
          "west": "side"
        },
        "id": 2099
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "6",
          "south": "side",
          "west": "none"
        },
        "id": 2100
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "6",
          "south": "none",
          "west": "up"
        },
        "id": 2101
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "6",
          "south": "none",
          "west": "side"
        },
        "id": 2102
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "6",
          "south": "none",
          "west": "none"
        },
        "id": 2103
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "7",
          "south": "up",
          "west": "up"
        },
        "id": 2104
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "7",
          "south": "up",
          "west": "side"
        },
        "id": 2105
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "7",
          "south": "up",
          "west": "none"
        },
        "id": 2106
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "7",
          "south": "side",
          "west": "up"
        },
        "id": 2107
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "7",
          "south": "side",
          "west": "side"
        },
        "id": 2108
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "7",
          "south": "side",
          "west": "none"
        },
        "id": 2109
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "7",
          "south": "none",
          "west": "up"
        },
        "id": 2110
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "7",
          "south": "none",
          "west": "side"
        },
        "id": 2111
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "7",
          "south": "none",
          "west": "none"
        },
        "id": 2112
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "8",
          "south": "up",
          "west": "up"
        },
        "id": 2113
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "8",
          "south": "up",
          "west": "side"
        },
        "id": 2114
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "8",
          "south": "up",
          "west": "none"
        },
        "id": 2115
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "8",
          "south": "side",
          "west": "up"
        },
        "id": 2116
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "8",
          "south": "side",
          "west": "side"
        },
        "id": 2117
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "8",
          "south": "side",
          "west": "none"
        },
        "id": 2118
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "8",
          "south": "none",
          "west": "up"
        },
        "id": 2119
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "8",
          "south": "none",
          "west": "side"
        },
        "id": 2120
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "8",
          "south": "none",
          "west": "none"
        },
        "id": 2121
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "9",
          "south": "up",
          "west": "up"
        },
        "id": 2122
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "9",
          "south": "up",
          "west": "side"
        },
        "id": 2123
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "9",
          "south": "up",
          "west": "none"
        },
        "id": 2124
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "9",
          "south": "side",
          "west": "up"
        },
        "id": 2125
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "9",
          "south": "side",
          "west": "side"
        },
        "id": 2126
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "9",
          "south": "side",
          "west": "none"
        },
        "id": 2127
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "9",
          "south": "none",
          "west": "up"
        },
        "id": 2128
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "9",
          "south": "none",
          "west": "side"
        },
        "id": 2129
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "9",
          "south": "none",
          "west": "none"
        },
        "id": 2130
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "10",
          "south": "up",
          "west": "up"
        },
        "id": 2131
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "10",
          "south": "up",
          "west": "side"
        },
        "id": 2132
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "10",
          "south": "up",
          "west": "none"
        },
        "id": 2133
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "10",
          "south": "side",
          "west": "up"
        },
        "id": 2134
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "10",
          "south": "side",
          "west": "side"
        },
        "id": 2135
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "10",
          "south": "side",
          "west": "none"
        },
        "id": 2136
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "10",
          "south": "none",
          "west": "up"
        },
        "id": 2137
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "10",
          "south": "none",
          "west": "side"
        },
        "id": 2138
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "10",
          "south": "none",
          "west": "none"
        },
        "id": 2139
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "11",
          "south": "up",
          "west": "up"
        },
        "id": 2140
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "11",
          "south": "up",
          "west": "side"
        },
        "id": 2141
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "11",
          "south": "up",
          "west": "none"
        },
        "id": 2142
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "11",
          "south": "side",
          "west": "up"
        },
        "id": 2143
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "11",
          "south": "side",
          "west": "side"
        },
        "id": 2144
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "11",
          "south": "side",
          "west": "none"
        },
        "id": 2145
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "11",
          "south": "none",
          "west": "up"
        },
        "id": 2146
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "11",
          "south": "none",
          "west": "side"
        },
        "id": 2147
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "11",
          "south": "none",
          "west": "none"
        },
        "id": 2148
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "12",
          "south": "up",
          "west": "up"
        },
        "id": 2149
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "12",
          "south": "up",
          "west": "side"
        },
        "id": 2150
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "12",
          "south": "up",
          "west": "none"
        },
        "id": 2151
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "12",
          "south": "side",
          "west": "up"
        },
        "id": 2152
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "12",
          "south": "side",
          "west": "side"
        },
        "id": 2153
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "12",
          "south": "side",
          "west": "none"
        },
        "id": 2154
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "12",
          "south": "none",
          "west": "up"
        },
        "id": 2155
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "12",
          "south": "none",
          "west": "side"
        },
        "id": 2156
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "12",
          "south": "none",
          "west": "none"
        },
        "id": 2157
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "13",
          "south": "up",
          "west": "up"
        },
        "id": 2158
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "13",
          "south": "up",
          "west": "side"
        },
        "id": 2159
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "13",
          "south": "up",
          "west": "none"
        },
        "id": 2160
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "13",
          "south": "side",
          "west": "up"
        },
        "id": 2161
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "13",
          "south": "side",
          "west": "side"
        },
        "id": 2162
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "13",
          "south": "side",
          "west": "none"
        },
        "id": 2163
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "13",
          "south": "none",
          "west": "up"
        },
        "id": 2164
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "13",
          "south": "none",
          "west": "side"
        },
        "id": 2165
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "13",
          "south": "none",
          "west": "none"
        },
        "id": 2166
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "14",
          "south": "up",
          "west": "up"
        },
        "id": 2167
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "14",
          "south": "up",
          "west": "side"
        },
        "id": 2168
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "14",
          "south": "up",
          "west": "none"
        },
        "id": 2169
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "14",
          "south": "side",
          "west": "up"
        },
        "id": 2170
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "14",
          "south": "side",
          "west": "side"
        },
        "id": 2171
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "14",
          "south": "side",
          "west": "none"
        },
        "id": 2172
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "14",
          "south": "none",
          "west": "up"
        },
        "id": 2173
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "14",
          "south": "none",
          "west": "side"
        },
        "id": 2174
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "14",
          "south": "none",
          "west": "none"
        },
        "id": 2175
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "15",
          "south": "up",
          "west": "up"
        },
        "id": 2176
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "15",
          "south": "up",
          "west": "side"
        },
        "id": 2177
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "15",
          "south": "up",
          "west": "none"
        },
        "id": 2178
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "15",
          "south": "side",
          "west": "up"
        },
        "id": 2179
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "15",
          "south": "side",
          "west": "side"
        },
        "id": 2180
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "15",
          "south": "side",
          "west": "none"
        },
        "id": 2181
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "15",
          "south": "none",
          "west": "up"
        },
        "id": 2182
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "15",
          "south": "none",
          "west": "side"
        },
        "id": 2183
      },
      {
        "properties": {
          "east": "up",
          "north": "none",
          "power": "15",
          "south": "none",
          "west": "none"
        },
        "id": 2184
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "0",
          "south": "up",
          "west": "up"
        },
        "id": 2185
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "0",
          "south": "up",
          "west": "side"
        },
        "id": 2186
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "0",
          "south": "up",
          "west": "none"
        },
        "id": 2187
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "0",
          "south": "side",
          "west": "up"
        },
        "id": 2188
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "0",
          "south": "side",
          "west": "side"
        },
        "id": 2189
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "0",
          "south": "side",
          "west": "none"
        },
        "id": 2190
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "0",
          "south": "none",
          "west": "up"
        },
        "id": 2191
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "0",
          "south": "none",
          "west": "side"
        },
        "id": 2192
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "0",
          "south": "none",
          "west": "none"
        },
        "id": 2193
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "1",
          "south": "up",
          "west": "up"
        },
        "id": 2194
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "1",
          "south": "up",
          "west": "side"
        },
        "id": 2195
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "1",
          "south": "up",
          "west": "none"
        },
        "id": 2196
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "1",
          "south": "side",
          "west": "up"
        },
        "id": 2197
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "1",
          "south": "side",
          "west": "side"
        },
        "id": 2198
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "1",
          "south": "side",
          "west": "none"
        },
        "id": 2199
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "1",
          "south": "none",
          "west": "up"
        },
        "id": 2200
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "1",
          "south": "none",
          "west": "side"
        },
        "id": 2201
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "1",
          "south": "none",
          "west": "none"
        },
        "id": 2202
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "2",
          "south": "up",
          "west": "up"
        },
        "id": 2203
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "2",
          "south": "up",
          "west": "side"
        },
        "id": 2204
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "2",
          "south": "up",
          "west": "none"
        },
        "id": 2205
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "2",
          "south": "side",
          "west": "up"
        },
        "id": 2206
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "2",
          "south": "side",
          "west": "side"
        },
        "id": 2207
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "2",
          "south": "side",
          "west": "none"
        },
        "id": 2208
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "2",
          "south": "none",
          "west": "up"
        },
        "id": 2209
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "2",
          "south": "none",
          "west": "side"
        },
        "id": 2210
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "2",
          "south": "none",
          "west": "none"
        },
        "id": 2211
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "3",
          "south": "up",
          "west": "up"
        },
        "id": 2212
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "3",
          "south": "up",
          "west": "side"
        },
        "id": 2213
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "3",
          "south": "up",
          "west": "none"
        },
        "id": 2214
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "3",
          "south": "side",
          "west": "up"
        },
        "id": 2215
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "3",
          "south": "side",
          "west": "side"
        },
        "id": 2216
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "3",
          "south": "side",
          "west": "none"
        },
        "id": 2217
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "3",
          "south": "none",
          "west": "up"
        },
        "id": 2218
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "3",
          "south": "none",
          "west": "side"
        },
        "id": 2219
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "3",
          "south": "none",
          "west": "none"
        },
        "id": 2220
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "4",
          "south": "up",
          "west": "up"
        },
        "id": 2221
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "4",
          "south": "up",
          "west": "side"
        },
        "id": 2222
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "4",
          "south": "up",
          "west": "none"
        },
        "id": 2223
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "4",
          "south": "side",
          "west": "up"
        },
        "id": 2224
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "4",
          "south": "side",
          "west": "side"
        },
        "id": 2225
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "4",
          "south": "side",
          "west": "none"
        },
        "id": 2226
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "4",
          "south": "none",
          "west": "up"
        },
        "id": 2227
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "4",
          "south": "none",
          "west": "side"
        },
        "id": 2228
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "4",
          "south": "none",
          "west": "none"
        },
        "id": 2229
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "5",
          "south": "up",
          "west": "up"
        },
        "id": 2230
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "5",
          "south": "up",
          "west": "side"
        },
        "id": 2231
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "5",
          "south": "up",
          "west": "none"
        },
        "id": 2232
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "5",
          "south": "side",
          "west": "up"
        },
        "id": 2233
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "5",
          "south": "side",
          "west": "side"
        },
        "id": 2234
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "5",
          "south": "side",
          "west": "none"
        },
        "id": 2235
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "5",
          "south": "none",
          "west": "up"
        },
        "id": 2236
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "5",
          "south": "none",
          "west": "side"
        },
        "id": 2237
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "5",
          "south": "none",
          "west": "none"
        },
        "id": 2238
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "6",
          "south": "up",
          "west": "up"
        },
        "id": 2239
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "6",
          "south": "up",
          "west": "side"
        },
        "id": 2240
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "6",
          "south": "up",
          "west": "none"
        },
        "id": 2241
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "6",
          "south": "side",
          "west": "up"
        },
        "id": 2242
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "6",
          "south": "side",
          "west": "side"
        },
        "id": 2243
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "6",
          "south": "side",
          "west": "none"
        },
        "id": 2244
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "6",
          "south": "none",
          "west": "up"
        },
        "id": 2245
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "6",
          "south": "none",
          "west": "side"
        },
        "id": 2246
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "6",
          "south": "none",
          "west": "none"
        },
        "id": 2247
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "7",
          "south": "up",
          "west": "up"
        },
        "id": 2248
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "7",
          "south": "up",
          "west": "side"
        },
        "id": 2249
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "7",
          "south": "up",
          "west": "none"
        },
        "id": 2250
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "7",
          "south": "side",
          "west": "up"
        },
        "id": 2251
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "7",
          "south": "side",
          "west": "side"
        },
        "id": 2252
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "7",
          "south": "side",
          "west": "none"
        },
        "id": 2253
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "7",
          "south": "none",
          "west": "up"
        },
        "id": 2254
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "7",
          "south": "none",
          "west": "side"
        },
        "id": 2255
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "7",
          "south": "none",
          "west": "none"
        },
        "id": 2256
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "8",
          "south": "up",
          "west": "up"
        },
        "id": 2257
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "8",
          "south": "up",
          "west": "side"
        },
        "id": 2258
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "8",
          "south": "up",
          "west": "none"
        },
        "id": 2259
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "8",
          "south": "side",
          "west": "up"
        },
        "id": 2260
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "8",
          "south": "side",
          "west": "side"
        },
        "id": 2261
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "8",
          "south": "side",
          "west": "none"
        },
        "id": 2262
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "8",
          "south": "none",
          "west": "up"
        },
        "id": 2263
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "8",
          "south": "none",
          "west": "side"
        },
        "id": 2264
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "8",
          "south": "none",
          "west": "none"
        },
        "id": 2265
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "9",
          "south": "up",
          "west": "up"
        },
        "id": 2266
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "9",
          "south": "up",
          "west": "side"
        },
        "id": 2267
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "9",
          "south": "up",
          "west": "none"
        },
        "id": 2268
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "9",
          "south": "side",
          "west": "up"
        },
        "id": 2269
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "9",
          "south": "side",
          "west": "side"
        },
        "id": 2270
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "9",
          "south": "side",
          "west": "none"
        },
        "id": 2271
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "9",
          "south": "none",
          "west": "up"
        },
        "id": 2272
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "9",
          "south": "none",
          "west": "side"
        },
        "id": 2273
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "9",
          "south": "none",
          "west": "none"
        },
        "id": 2274
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "10",
          "south": "up",
          "west": "up"
        },
        "id": 2275
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "10",
          "south": "up",
          "west": "side"
        },
        "id": 2276
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "10",
          "south": "up",
          "west": "none"
        },
        "id": 2277
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "10",
          "south": "side",
          "west": "up"
        },
        "id": 2278
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "10",
          "south": "side",
          "west": "side"
        },
        "id": 2279
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "10",
          "south": "side",
          "west": "none"
        },
        "id": 2280
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "10",
          "south": "none",
          "west": "up"
        },
        "id": 2281
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "10",
          "south": "none",
          "west": "side"
        },
        "id": 2282
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "10",
          "south": "none",
          "west": "none"
        },
        "id": 2283
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "11",
          "south": "up",
          "west": "up"
        },
        "id": 2284
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "11",
          "south": "up",
          "west": "side"
        },
        "id": 2285
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "11",
          "south": "up",
          "west": "none"
        },
        "id": 2286
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "11",
          "south": "side",
          "west": "up"
        },
        "id": 2287
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "11",
          "south": "side",
          "west": "side"
        },
        "id": 2288
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "11",
          "south": "side",
          "west": "none"
        },
        "id": 2289
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "11",
          "south": "none",
          "west": "up"
        },
        "id": 2290
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "11",
          "south": "none",
          "west": "side"
        },
        "id": 2291
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "11",
          "south": "none",
          "west": "none"
        },
        "id": 2292
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "12",
          "south": "up",
          "west": "up"
        },
        "id": 2293
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "12",
          "south": "up",
          "west": "side"
        },
        "id": 2294
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "12",
          "south": "up",
          "west": "none"
        },
        "id": 2295
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "12",
          "south": "side",
          "west": "up"
        },
        "id": 2296
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "12",
          "south": "side",
          "west": "side"
        },
        "id": 2297
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "12",
          "south": "side",
          "west": "none"
        },
        "id": 2298
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "12",
          "south": "none",
          "west": "up"
        },
        "id": 2299
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "12",
          "south": "none",
          "west": "side"
        },
        "id": 2300
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "12",
          "south": "none",
          "west": "none"
        },
        "id": 2301
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "13",
          "south": "up",
          "west": "up"
        },
        "id": 2302
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "13",
          "south": "up",
          "west": "side"
        },
        "id": 2303
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "13",
          "south": "up",
          "west": "none"
        },
        "id": 2304
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "13",
          "south": "side",
          "west": "up"
        },
        "id": 2305
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "13",
          "south": "side",
          "west": "side"
        },
        "id": 2306
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "13",
          "south": "side",
          "west": "none"
        },
        "id": 2307
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "13",
          "south": "none",
          "west": "up"
        },
        "id": 2308
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "13",
          "south": "none",
          "west": "side"
        },
        "id": 2309
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "13",
          "south": "none",
          "west": "none"
        },
        "id": 2310
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "14",
          "south": "up",
          "west": "up"
        },
        "id": 2311
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "14",
          "south": "up",
          "west": "side"
        },
        "id": 2312
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "14",
          "south": "up",
          "west": "none"
        },
        "id": 2313
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "14",
          "south": "side",
          "west": "up"
        },
        "id": 2314
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "14",
          "south": "side",
          "west": "side"
        },
        "id": 2315
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "14",
          "south": "side",
          "west": "none"
        },
        "id": 2316
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "14",
          "south": "none",
          "west": "up"
        },
        "id": 2317
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "14",
          "south": "none",
          "west": "side"
        },
        "id": 2318
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "14",
          "south": "none",
          "west": "none"
        },
        "id": 2319
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "15",
          "south": "up",
          "west": "up"
        },
        "id": 2320
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "15",
          "south": "up",
          "west": "side"
        },
        "id": 2321
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "15",
          "south": "up",
          "west": "none"
        },
        "id": 2322
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "15",
          "south": "side",
          "west": "up"
        },
        "id": 2323
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "15",
          "south": "side",
          "west": "side"
        },
        "id": 2324
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "15",
          "south": "side",
          "west": "none"
        },
        "id": 2325
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "15",
          "south": "none",
          "west": "up"
        },
        "id": 2326
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "15",
          "south": "none",
          "west": "side"
        },
        "id": 2327
      },
      {
        "properties": {
          "east": "side",
          "north": "up",
          "power": "15",
          "south": "none",
          "west": "none"
        },
        "id": 2328
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "0",
          "south": "up",
          "west": "up"
        },
        "id": 2329
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "0",
          "south": "up",
          "west": "side"
        },
        "id": 2330
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "0",
          "south": "up",
          "west": "none"
        },
        "id": 2331
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "0",
          "south": "side",
          "west": "up"
        },
        "id": 2332
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "0",
          "south": "side",
          "west": "side"
        },
        "id": 2333
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "0",
          "south": "side",
          "west": "none"
        },
        "id": 2334
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "0",
          "south": "none",
          "west": "up"
        },
        "id": 2335
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "0",
          "south": "none",
          "west": "side"
        },
        "id": 2336
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "0",
          "south": "none",
          "west": "none"
        },
        "id": 2337
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "1",
          "south": "up",
          "west": "up"
        },
        "id": 2338
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "1",
          "south": "up",
          "west": "side"
        },
        "id": 2339
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "1",
          "south": "up",
          "west": "none"
        },
        "id": 2340
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "1",
          "south": "side",
          "west": "up"
        },
        "id": 2341
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "1",
          "south": "side",
          "west": "side"
        },
        "id": 2342
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "1",
          "south": "side",
          "west": "none"
        },
        "id": 2343
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "1",
          "south": "none",
          "west": "up"
        },
        "id": 2344
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "1",
          "south": "none",
          "west": "side"
        },
        "id": 2345
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "1",
          "south": "none",
          "west": "none"
        },
        "id": 2346
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "2",
          "south": "up",
          "west": "up"
        },
        "id": 2347
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "2",
          "south": "up",
          "west": "side"
        },
        "id": 2348
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "2",
          "south": "up",
          "west": "none"
        },
        "id": 2349
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "2",
          "south": "side",
          "west": "up"
        },
        "id": 2350
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "2",
          "south": "side",
          "west": "side"
        },
        "id": 2351
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "2",
          "south": "side",
          "west": "none"
        },
        "id": 2352
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "2",
          "south": "none",
          "west": "up"
        },
        "id": 2353
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "2",
          "south": "none",
          "west": "side"
        },
        "id": 2354
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "2",
          "south": "none",
          "west": "none"
        },
        "id": 2355
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "3",
          "south": "up",
          "west": "up"
        },
        "id": 2356
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "3",
          "south": "up",
          "west": "side"
        },
        "id": 2357
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "3",
          "south": "up",
          "west": "none"
        },
        "id": 2358
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "3",
          "south": "side",
          "west": "up"
        },
        "id": 2359
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "3",
          "south": "side",
          "west": "side"
        },
        "id": 2360
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "3",
          "south": "side",
          "west": "none"
        },
        "id": 2361
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "3",
          "south": "none",
          "west": "up"
        },
        "id": 2362
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "3",
          "south": "none",
          "west": "side"
        },
        "id": 2363
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "3",
          "south": "none",
          "west": "none"
        },
        "id": 2364
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "4",
          "south": "up",
          "west": "up"
        },
        "id": 2365
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "4",
          "south": "up",
          "west": "side"
        },
        "id": 2366
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "4",
          "south": "up",
          "west": "none"
        },
        "id": 2367
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "4",
          "south": "side",
          "west": "up"
        },
        "id": 2368
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "4",
          "south": "side",
          "west": "side"
        },
        "id": 2369
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "4",
          "south": "side",
          "west": "none"
        },
        "id": 2370
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "4",
          "south": "none",
          "west": "up"
        },
        "id": 2371
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "4",
          "south": "none",
          "west": "side"
        },
        "id": 2372
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "4",
          "south": "none",
          "west": "none"
        },
        "id": 2373
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "5",
          "south": "up",
          "west": "up"
        },
        "id": 2374
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "5",
          "south": "up",
          "west": "side"
        },
        "id": 2375
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "5",
          "south": "up",
          "west": "none"
        },
        "id": 2376
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "5",
          "south": "side",
          "west": "up"
        },
        "id": 2377
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "5",
          "south": "side",
          "west": "side"
        },
        "id": 2378
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "5",
          "south": "side",
          "west": "none"
        },
        "id": 2379
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "5",
          "south": "none",
          "west": "up"
        },
        "id": 2380
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "5",
          "south": "none",
          "west": "side"
        },
        "id": 2381
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "5",
          "south": "none",
          "west": "none"
        },
        "id": 2382
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "6",
          "south": "up",
          "west": "up"
        },
        "id": 2383
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "6",
          "south": "up",
          "west": "side"
        },
        "id": 2384
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "6",
          "south": "up",
          "west": "none"
        },
        "id": 2385
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "6",
          "south": "side",
          "west": "up"
        },
        "id": 2386
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "6",
          "south": "side",
          "west": "side"
        },
        "id": 2387
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "6",
          "south": "side",
          "west": "none"
        },
        "id": 2388
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "6",
          "south": "none",
          "west": "up"
        },
        "id": 2389
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "6",
          "south": "none",
          "west": "side"
        },
        "id": 2390
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "6",
          "south": "none",
          "west": "none"
        },
        "id": 2391
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "7",
          "south": "up",
          "west": "up"
        },
        "id": 2392
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "7",
          "south": "up",
          "west": "side"
        },
        "id": 2393
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "7",
          "south": "up",
          "west": "none"
        },
        "id": 2394
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "7",
          "south": "side",
          "west": "up"
        },
        "id": 2395
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "7",
          "south": "side",
          "west": "side"
        },
        "id": 2396
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "7",
          "south": "side",
          "west": "none"
        },
        "id": 2397
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "7",
          "south": "none",
          "west": "up"
        },
        "id": 2398
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "7",
          "south": "none",
          "west": "side"
        },
        "id": 2399
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "7",
          "south": "none",
          "west": "none"
        },
        "id": 2400
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "8",
          "south": "up",
          "west": "up"
        },
        "id": 2401
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "8",
          "south": "up",
          "west": "side"
        },
        "id": 2402
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "8",
          "south": "up",
          "west": "none"
        },
        "id": 2403
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "8",
          "south": "side",
          "west": "up"
        },
        "id": 2404
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "8",
          "south": "side",
          "west": "side"
        },
        "id": 2405
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "8",
          "south": "side",
          "west": "none"
        },
        "id": 2406
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "8",
          "south": "none",
          "west": "up"
        },
        "id": 2407
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "8",
          "south": "none",
          "west": "side"
        },
        "id": 2408
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "8",
          "south": "none",
          "west": "none"
        },
        "id": 2409
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "9",
          "south": "up",
          "west": "up"
        },
        "id": 2410
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "9",
          "south": "up",
          "west": "side"
        },
        "id": 2411
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "9",
          "south": "up",
          "west": "none"
        },
        "id": 2412
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "9",
          "south": "side",
          "west": "up"
        },
        "id": 2413
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "9",
          "south": "side",
          "west": "side"
        },
        "id": 2414
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "9",
          "south": "side",
          "west": "none"
        },
        "id": 2415
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "9",
          "south": "none",
          "west": "up"
        },
        "id": 2416
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "9",
          "south": "none",
          "west": "side"
        },
        "id": 2417
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "9",
          "south": "none",
          "west": "none"
        },
        "id": 2418
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "10",
          "south": "up",
          "west": "up"
        },
        "id": 2419
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "10",
          "south": "up",
          "west": "side"
        },
        "id": 2420
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "10",
          "south": "up",
          "west": "none"
        },
        "id": 2421
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "10",
          "south": "side",
          "west": "up"
        },
        "id": 2422
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "10",
          "south": "side",
          "west": "side"
        },
        "id": 2423
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "10",
          "south": "side",
          "west": "none"
        },
        "id": 2424
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "10",
          "south": "none",
          "west": "up"
        },
        "id": 2425
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "10",
          "south": "none",
          "west": "side"
        },
        "id": 2426
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "10",
          "south": "none",
          "west": "none"
        },
        "id": 2427
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "11",
          "south": "up",
          "west": "up"
        },
        "id": 2428
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "11",
          "south": "up",
          "west": "side"
        },
        "id": 2429
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "11",
          "south": "up",
          "west": "none"
        },
        "id": 2430
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "11",
          "south": "side",
          "west": "up"
        },
        "id": 2431
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "11",
          "south": "side",
          "west": "side"
        },
        "id": 2432
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "11",
          "south": "side",
          "west": "none"
        },
        "id": 2433
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "11",
          "south": "none",
          "west": "up"
        },
        "id": 2434
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "11",
          "south": "none",
          "west": "side"
        },
        "id": 2435
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "11",
          "south": "none",
          "west": "none"
        },
        "id": 2436
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "12",
          "south": "up",
          "west": "up"
        },
        "id": 2437
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "12",
          "south": "up",
          "west": "side"
        },
        "id": 2438
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "12",
          "south": "up",
          "west": "none"
        },
        "id": 2439
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "12",
          "south": "side",
          "west": "up"
        },
        "id": 2440
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "12",
          "south": "side",
          "west": "side"
        },
        "id": 2441
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "12",
          "south": "side",
          "west": "none"
        },
        "id": 2442
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "12",
          "south": "none",
          "west": "up"
        },
        "id": 2443
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "12",
          "south": "none",
          "west": "side"
        },
        "id": 2444
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "12",
          "south": "none",
          "west": "none"
        },
        "id": 2445
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "13",
          "south": "up",
          "west": "up"
        },
        "id": 2446
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "13",
          "south": "up",
          "west": "side"
        },
        "id": 2447
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "13",
          "south": "up",
          "west": "none"
        },
        "id": 2448
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "13",
          "south": "side",
          "west": "up"
        },
        "id": 2449
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "13",
          "south": "side",
          "west": "side"
        },
        "id": 2450
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "13",
          "south": "side",
          "west": "none"
        },
        "id": 2451
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "13",
          "south": "none",
          "west": "up"
        },
        "id": 2452
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "13",
          "south": "none",
          "west": "side"
        },
        "id": 2453
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "13",
          "south": "none",
          "west": "none"
        },
        "id": 2454
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "14",
          "south": "up",
          "west": "up"
        },
        "id": 2455
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "14",
          "south": "up",
          "west": "side"
        },
        "id": 2456
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "14",
          "south": "up",
          "west": "none"
        },
        "id": 2457
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "14",
          "south": "side",
          "west": "up"
        },
        "id": 2458
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "14",
          "south": "side",
          "west": "side"
        },
        "id": 2459
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "14",
          "south": "side",
          "west": "none"
        },
        "id": 2460
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "14",
          "south": "none",
          "west": "up"
        },
        "id": 2461
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "14",
          "south": "none",
          "west": "side"
        },
        "id": 2462
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "14",
          "south": "none",
          "west": "none"
        },
        "id": 2463
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "15",
          "south": "up",
          "west": "up"
        },
        "id": 2464
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "15",
          "south": "up",
          "west": "side"
        },
        "id": 2465
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "15",
          "south": "up",
          "west": "none"
        },
        "id": 2466
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "15",
          "south": "side",
          "west": "up"
        },
        "id": 2467
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "15",
          "south": "side",
          "west": "side"
        },
        "id": 2468
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "15",
          "south": "side",
          "west": "none"
        },
        "id": 2469
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "15",
          "south": "none",
          "west": "up"
        },
        "id": 2470
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "15",
          "south": "none",
          "west": "side"
        },
        "id": 2471
      },
      {
        "properties": {
          "east": "side",
          "north": "side",
          "power": "15",
          "south": "none",
          "west": "none"
        },
        "id": 2472
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "0",
          "south": "up",
          "west": "up"
        },
        "id": 2473
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "0",
          "south": "up",
          "west": "side"
        },
        "id": 2474
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "0",
          "south": "up",
          "west": "none"
        },
        "id": 2475
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "0",
          "south": "side",
          "west": "up"
        },
        "id": 2476
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "0",
          "south": "side",
          "west": "side"
        },
        "id": 2477
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "0",
          "south": "side",
          "west": "none"
        },
        "id": 2478
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "0",
          "south": "none",
          "west": "up"
        },
        "id": 2479
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "0",
          "south": "none",
          "west": "side"
        },
        "id": 2480
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "0",
          "south": "none",
          "west": "none"
        },
        "id": 2481
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "1",
          "south": "up",
          "west": "up"
        },
        "id": 2482
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "1",
          "south": "up",
          "west": "side"
        },
        "id": 2483
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "1",
          "south": "up",
          "west": "none"
        },
        "id": 2484
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "1",
          "south": "side",
          "west": "up"
        },
        "id": 2485
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "1",
          "south": "side",
          "west": "side"
        },
        "id": 2486
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "1",
          "south": "side",
          "west": "none"
        },
        "id": 2487
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "1",
          "south": "none",
          "west": "up"
        },
        "id": 2488
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "1",
          "south": "none",
          "west": "side"
        },
        "id": 2489
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "1",
          "south": "none",
          "west": "none"
        },
        "id": 2490
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "2",
          "south": "up",
          "west": "up"
        },
        "id": 2491
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "2",
          "south": "up",
          "west": "side"
        },
        "id": 2492
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "2",
          "south": "up",
          "west": "none"
        },
        "id": 2493
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "2",
          "south": "side",
          "west": "up"
        },
        "id": 2494
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "2",
          "south": "side",
          "west": "side"
        },
        "id": 2495
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "2",
          "south": "side",
          "west": "none"
        },
        "id": 2496
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "2",
          "south": "none",
          "west": "up"
        },
        "id": 2497
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "2",
          "south": "none",
          "west": "side"
        },
        "id": 2498
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "2",
          "south": "none",
          "west": "none"
        },
        "id": 2499
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "3",
          "south": "up",
          "west": "up"
        },
        "id": 2500
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "3",
          "south": "up",
          "west": "side"
        },
        "id": 2501
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "3",
          "south": "up",
          "west": "none"
        },
        "id": 2502
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "3",
          "south": "side",
          "west": "up"
        },
        "id": 2503
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "3",
          "south": "side",
          "west": "side"
        },
        "id": 2504
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "3",
          "south": "side",
          "west": "none"
        },
        "id": 2505
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "3",
          "south": "none",
          "west": "up"
        },
        "id": 2506
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "3",
          "south": "none",
          "west": "side"
        },
        "id": 2507
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "3",
          "south": "none",
          "west": "none"
        },
        "id": 2508
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "4",
          "south": "up",
          "west": "up"
        },
        "id": 2509
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "4",
          "south": "up",
          "west": "side"
        },
        "id": 2510
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "4",
          "south": "up",
          "west": "none"
        },
        "id": 2511
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "4",
          "south": "side",
          "west": "up"
        },
        "id": 2512
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "4",
          "south": "side",
          "west": "side"
        },
        "id": 2513
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "4",
          "south": "side",
          "west": "none"
        },
        "id": 2514
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "4",
          "south": "none",
          "west": "up"
        },
        "id": 2515
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "4",
          "south": "none",
          "west": "side"
        },
        "id": 2516
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "4",
          "south": "none",
          "west": "none"
        },
        "id": 2517
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "5",
          "south": "up",
          "west": "up"
        },
        "id": 2518
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "5",
          "south": "up",
          "west": "side"
        },
        "id": 2519
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "5",
          "south": "up",
          "west": "none"
        },
        "id": 2520
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "5",
          "south": "side",
          "west": "up"
        },
        "id": 2521
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "5",
          "south": "side",
          "west": "side"
        },
        "id": 2522
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "5",
          "south": "side",
          "west": "none"
        },
        "id": 2523
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "5",
          "south": "none",
          "west": "up"
        },
        "id": 2524
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "5",
          "south": "none",
          "west": "side"
        },
        "id": 2525
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "5",
          "south": "none",
          "west": "none"
        },
        "id": 2526
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "6",
          "south": "up",
          "west": "up"
        },
        "id": 2527
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "6",
          "south": "up",
          "west": "side"
        },
        "id": 2528
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "6",
          "south": "up",
          "west": "none"
        },
        "id": 2529
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "6",
          "south": "side",
          "west": "up"
        },
        "id": 2530
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "6",
          "south": "side",
          "west": "side"
        },
        "id": 2531
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "6",
          "south": "side",
          "west": "none"
        },
        "id": 2532
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "6",
          "south": "none",
          "west": "up"
        },
        "id": 2533
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "6",
          "south": "none",
          "west": "side"
        },
        "id": 2534
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "6",
          "south": "none",
          "west": "none"
        },
        "id": 2535
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "7",
          "south": "up",
          "west": "up"
        },
        "id": 2536
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "7",
          "south": "up",
          "west": "side"
        },
        "id": 2537
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "7",
          "south": "up",
          "west": "none"
        },
        "id": 2538
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "7",
          "south": "side",
          "west": "up"
        },
        "id": 2539
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "7",
          "south": "side",
          "west": "side"
        },
        "id": 2540
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "7",
          "south": "side",
          "west": "none"
        },
        "id": 2541
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "7",
          "south": "none",
          "west": "up"
        },
        "id": 2542
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "7",
          "south": "none",
          "west": "side"
        },
        "id": 2543
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "7",
          "south": "none",
          "west": "none"
        },
        "id": 2544
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "8",
          "south": "up",
          "west": "up"
        },
        "id": 2545
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "8",
          "south": "up",
          "west": "side"
        },
        "id": 2546
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "8",
          "south": "up",
          "west": "none"
        },
        "id": 2547
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "8",
          "south": "side",
          "west": "up"
        },
        "id": 2548
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "8",
          "south": "side",
          "west": "side"
        },
        "id": 2549
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "8",
          "south": "side",
          "west": "none"
        },
        "id": 2550
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "8",
          "south": "none",
          "west": "up"
        },
        "id": 2551
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "8",
          "south": "none",
          "west": "side"
        },
        "id": 2552
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "8",
          "south": "none",
          "west": "none"
        },
        "id": 2553
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "9",
          "south": "up",
          "west": "up"
        },
        "id": 2554
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "9",
          "south": "up",
          "west": "side"
        },
        "id": 2555
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "9",
          "south": "up",
          "west": "none"
        },
        "id": 2556
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "9",
          "south": "side",
          "west": "up"
        },
        "id": 2557
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "9",
          "south": "side",
          "west": "side"
        },
        "id": 2558
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "9",
          "south": "side",
          "west": "none"
        },
        "id": 2559
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "9",
          "south": "none",
          "west": "up"
        },
        "id": 2560
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "9",
          "south": "none",
          "west": "side"
        },
        "id": 2561
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "9",
          "south": "none",
          "west": "none"
        },
        "id": 2562
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "10",
          "south": "up",
          "west": "up"
        },
        "id": 2563
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "10",
          "south": "up",
          "west": "side"
        },
        "id": 2564
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "10",
          "south": "up",
          "west": "none"
        },
        "id": 2565
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "10",
          "south": "side",
          "west": "up"
        },
        "id": 2566
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "10",
          "south": "side",
          "west": "side"
        },
        "id": 2567
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "10",
          "south": "side",
          "west": "none"
        },
        "id": 2568
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "10",
          "south": "none",
          "west": "up"
        },
        "id": 2569
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "10",
          "south": "none",
          "west": "side"
        },
        "id": 2570
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "10",
          "south": "none",
          "west": "none"
        },
        "id": 2571
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "11",
          "south": "up",
          "west": "up"
        },
        "id": 2572
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "11",
          "south": "up",
          "west": "side"
        },
        "id": 2573
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "11",
          "south": "up",
          "west": "none"
        },
        "id": 2574
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "11",
          "south": "side",
          "west": "up"
        },
        "id": 2575
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "11",
          "south": "side",
          "west": "side"
        },
        "id": 2576
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "11",
          "south": "side",
          "west": "none"
        },
        "id": 2577
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "11",
          "south": "none",
          "west": "up"
        },
        "id": 2578
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "11",
          "south": "none",
          "west": "side"
        },
        "id": 2579
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "11",
          "south": "none",
          "west": "none"
        },
        "id": 2580
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "12",
          "south": "up",
          "west": "up"
        },
        "id": 2581
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "12",
          "south": "up",
          "west": "side"
        },
        "id": 2582
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "12",
          "south": "up",
          "west": "none"
        },
        "id": 2583
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "12",
          "south": "side",
          "west": "up"
        },
        "id": 2584
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "12",
          "south": "side",
          "west": "side"
        },
        "id": 2585
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "12",
          "south": "side",
          "west": "none"
        },
        "id": 2586
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "12",
          "south": "none",
          "west": "up"
        },
        "id": 2587
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "12",
          "south": "none",
          "west": "side"
        },
        "id": 2588
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "12",
          "south": "none",
          "west": "none"
        },
        "id": 2589
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "13",
          "south": "up",
          "west": "up"
        },
        "id": 2590
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "13",
          "south": "up",
          "west": "side"
        },
        "id": 2591
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "13",
          "south": "up",
          "west": "none"
        },
        "id": 2592
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "13",
          "south": "side",
          "west": "up"
        },
        "id": 2593
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "13",
          "south": "side",
          "west": "side"
        },
        "id": 2594
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "13",
          "south": "side",
          "west": "none"
        },
        "id": 2595
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "13",
          "south": "none",
          "west": "up"
        },
        "id": 2596
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "13",
          "south": "none",
          "west": "side"
        },
        "id": 2597
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "13",
          "south": "none",
          "west": "none"
        },
        "id": 2598
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "14",
          "south": "up",
          "west": "up"
        },
        "id": 2599
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "14",
          "south": "up",
          "west": "side"
        },
        "id": 2600
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "14",
          "south": "up",
          "west": "none"
        },
        "id": 2601
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "14",
          "south": "side",
          "west": "up"
        },
        "id": 2602
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "14",
          "south": "side",
          "west": "side"
        },
        "id": 2603
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "14",
          "south": "side",
          "west": "none"
        },
        "id": 2604
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "14",
          "south": "none",
          "west": "up"
        },
        "id": 2605
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "14",
          "south": "none",
          "west": "side"
        },
        "id": 2606
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "14",
          "south": "none",
          "west": "none"
        },
        "id": 2607
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "15",
          "south": "up",
          "west": "up"
        },
        "id": 2608
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "15",
          "south": "up",
          "west": "side"
        },
        "id": 2609
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "15",
          "south": "up",
          "west": "none"
        },
        "id": 2610
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "15",
          "south": "side",
          "west": "up"
        },
        "id": 2611
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "15",
          "south": "side",
          "west": "side"
        },
        "id": 2612
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "15",
          "south": "side",
          "west": "none"
        },
        "id": 2613
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "15",
          "south": "none",
          "west": "up"
        },
        "id": 2614
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "15",
          "south": "none",
          "west": "side"
        },
        "id": 2615
      },
      {
        "properties": {
          "east": "side",
          "north": "none",
          "power": "15",
          "south": "none",
          "west": "none"
        },
        "id": 2616
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "0",
          "south": "up",
          "west": "up"
        },
        "id": 2617
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "0",
          "south": "up",
          "west": "side"
        },
        "id": 2618
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "0",
          "south": "up",
          "west": "none"
        },
        "id": 2619
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "0",
          "south": "side",
          "west": "up"
        },
        "id": 2620
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "0",
          "south": "side",
          "west": "side"
        },
        "id": 2621
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "0",
          "south": "side",
          "west": "none"
        },
        "id": 2622
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "0",
          "south": "none",
          "west": "up"
        },
        "id": 2623
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "0",
          "south": "none",
          "west": "side"
        },
        "id": 2624
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "0",
          "south": "none",
          "west": "none"
        },
        "id": 2625
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "1",
          "south": "up",
          "west": "up"
        },
        "id": 2626
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "1",
          "south": "up",
          "west": "side"
        },
        "id": 2627
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "1",
          "south": "up",
          "west": "none"
        },
        "id": 2628
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "1",
          "south": "side",
          "west": "up"
        },
        "id": 2629
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "1",
          "south": "side",
          "west": "side"
        },
        "id": 2630
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "1",
          "south": "side",
          "west": "none"
        },
        "id": 2631
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "1",
          "south": "none",
          "west": "up"
        },
        "id": 2632
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "1",
          "south": "none",
          "west": "side"
        },
        "id": 2633
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "1",
          "south": "none",
          "west": "none"
        },
        "id": 2634
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "2",
          "south": "up",
          "west": "up"
        },
        "id": 2635
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "2",
          "south": "up",
          "west": "side"
        },
        "id": 2636
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "2",
          "south": "up",
          "west": "none"
        },
        "id": 2637
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "2",
          "south": "side",
          "west": "up"
        },
        "id": 2638
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "2",
          "south": "side",
          "west": "side"
        },
        "id": 2639
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "2",
          "south": "side",
          "west": "none"
        },
        "id": 2640
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "2",
          "south": "none",
          "west": "up"
        },
        "id": 2641
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "2",
          "south": "none",
          "west": "side"
        },
        "id": 2642
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "2",
          "south": "none",
          "west": "none"
        },
        "id": 2643
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "3",
          "south": "up",
          "west": "up"
        },
        "id": 2644
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "3",
          "south": "up",
          "west": "side"
        },
        "id": 2645
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "3",
          "south": "up",
          "west": "none"
        },
        "id": 2646
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "3",
          "south": "side",
          "west": "up"
        },
        "id": 2647
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "3",
          "south": "side",
          "west": "side"
        },
        "id": 2648
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "3",
          "south": "side",
          "west": "none"
        },
        "id": 2649
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "3",
          "south": "none",
          "west": "up"
        },
        "id": 2650
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "3",
          "south": "none",
          "west": "side"
        },
        "id": 2651
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "3",
          "south": "none",
          "west": "none"
        },
        "id": 2652
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "4",
          "south": "up",
          "west": "up"
        },
        "id": 2653
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "4",
          "south": "up",
          "west": "side"
        },
        "id": 2654
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "4",
          "south": "up",
          "west": "none"
        },
        "id": 2655
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "4",
          "south": "side",
          "west": "up"
        },
        "id": 2656
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "4",
          "south": "side",
          "west": "side"
        },
        "id": 2657
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "4",
          "south": "side",
          "west": "none"
        },
        "id": 2658
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "4",
          "south": "none",
          "west": "up"
        },
        "id": 2659
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "4",
          "south": "none",
          "west": "side"
        },
        "id": 2660
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "4",
          "south": "none",
          "west": "none"
        },
        "id": 2661
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "5",
          "south": "up",
          "west": "up"
        },
        "id": 2662
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "5",
          "south": "up",
          "west": "side"
        },
        "id": 2663
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "5",
          "south": "up",
          "west": "none"
        },
        "id": 2664
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "5",
          "south": "side",
          "west": "up"
        },
        "id": 2665
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "5",
          "south": "side",
          "west": "side"
        },
        "id": 2666
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "5",
          "south": "side",
          "west": "none"
        },
        "id": 2667
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "5",
          "south": "none",
          "west": "up"
        },
        "id": 2668
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "5",
          "south": "none",
          "west": "side"
        },
        "id": 2669
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "5",
          "south": "none",
          "west": "none"
        },
        "id": 2670
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "6",
          "south": "up",
          "west": "up"
        },
        "id": 2671
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "6",
          "south": "up",
          "west": "side"
        },
        "id": 2672
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "6",
          "south": "up",
          "west": "none"
        },
        "id": 2673
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "6",
          "south": "side",
          "west": "up"
        },
        "id": 2674
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "6",
          "south": "side",
          "west": "side"
        },
        "id": 2675
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "6",
          "south": "side",
          "west": "none"
        },
        "id": 2676
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "6",
          "south": "none",
          "west": "up"
        },
        "id": 2677
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "6",
          "south": "none",
          "west": "side"
        },
        "id": 2678
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "6",
          "south": "none",
          "west": "none"
        },
        "id": 2679
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "7",
          "south": "up",
          "west": "up"
        },
        "id": 2680
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "7",
          "south": "up",
          "west": "side"
        },
        "id": 2681
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "7",
          "south": "up",
          "west": "none"
        },
        "id": 2682
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "7",
          "south": "side",
          "west": "up"
        },
        "id": 2683
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "7",
          "south": "side",
          "west": "side"
        },
        "id": 2684
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "7",
          "south": "side",
          "west": "none"
        },
        "id": 2685
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "7",
          "south": "none",
          "west": "up"
        },
        "id": 2686
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "7",
          "south": "none",
          "west": "side"
        },
        "id": 2687
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "7",
          "south": "none",
          "west": "none"
        },
        "id": 2688
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "8",
          "south": "up",
          "west": "up"
        },
        "id": 2689
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "8",
          "south": "up",
          "west": "side"
        },
        "id": 2690
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "8",
          "south": "up",
          "west": "none"
        },
        "id": 2691
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "8",
          "south": "side",
          "west": "up"
        },
        "id": 2692
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "8",
          "south": "side",
          "west": "side"
        },
        "id": 2693
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "8",
          "south": "side",
          "west": "none"
        },
        "id": 2694
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "8",
          "south": "none",
          "west": "up"
        },
        "id": 2695
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "8",
          "south": "none",
          "west": "side"
        },
        "id": 2696
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "8",
          "south": "none",
          "west": "none"
        },
        "id": 2697
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "9",
          "south": "up",
          "west": "up"
        },
        "id": 2698
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "9",
          "south": "up",
          "west": "side"
        },
        "id": 2699
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "9",
          "south": "up",
          "west": "none"
        },
        "id": 2700
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "9",
          "south": "side",
          "west": "up"
        },
        "id": 2701
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "9",
          "south": "side",
          "west": "side"
        },
        "id": 2702
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "9",
          "south": "side",
          "west": "none"
        },
        "id": 2703
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "9",
          "south": "none",
          "west": "up"
        },
        "id": 2704
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "9",
          "south": "none",
          "west": "side"
        },
        "id": 2705
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "9",
          "south": "none",
          "west": "none"
        },
        "id": 2706
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "10",
          "south": "up",
          "west": "up"
        },
        "id": 2707
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "10",
          "south": "up",
          "west": "side"
        },
        "id": 2708
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "10",
          "south": "up",
          "west": "none"
        },
        "id": 2709
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "10",
          "south": "side",
          "west": "up"
        },
        "id": 2710
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "10",
          "south": "side",
          "west": "side"
        },
        "id": 2711
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "10",
          "south": "side",
          "west": "none"
        },
        "id": 2712
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "10",
          "south": "none",
          "west": "up"
        },
        "id": 2713
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "10",
          "south": "none",
          "west": "side"
        },
        "id": 2714
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "10",
          "south": "none",
          "west": "none"
        },
        "id": 2715
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "11",
          "south": "up",
          "west": "up"
        },
        "id": 2716
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "11",
          "south": "up",
          "west": "side"
        },
        "id": 2717
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "11",
          "south": "up",
          "west": "none"
        },
        "id": 2718
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "11",
          "south": "side",
          "west": "up"
        },
        "id": 2719
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "11",
          "south": "side",
          "west": "side"
        },
        "id": 2720
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "11",
          "south": "side",
          "west": "none"
        },
        "id": 2721
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "11",
          "south": "none",
          "west": "up"
        },
        "id": 2722
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "11",
          "south": "none",
          "west": "side"
        },
        "id": 2723
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "11",
          "south": "none",
          "west": "none"
        },
        "id": 2724
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "12",
          "south": "up",
          "west": "up"
        },
        "id": 2725
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "12",
          "south": "up",
          "west": "side"
        },
        "id": 2726
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "12",
          "south": "up",
          "west": "none"
        },
        "id": 2727
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "12",
          "south": "side",
          "west": "up"
        },
        "id": 2728
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "12",
          "south": "side",
          "west": "side"
        },
        "id": 2729
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "12",
          "south": "side",
          "west": "none"
        },
        "id": 2730
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "12",
          "south": "none",
          "west": "up"
        },
        "id": 2731
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "12",
          "south": "none",
          "west": "side"
        },
        "id": 2732
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "12",
          "south": "none",
          "west": "none"
        },
        "id": 2733
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "13",
          "south": "up",
          "west": "up"
        },
        "id": 2734
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "13",
          "south": "up",
          "west": "side"
        },
        "id": 2735
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "13",
          "south": "up",
          "west": "none"
        },
        "id": 2736
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "13",
          "south": "side",
          "west": "up"
        },
        "id": 2737
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "13",
          "south": "side",
          "west": "side"
        },
        "id": 2738
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "13",
          "south": "side",
          "west": "none"
        },
        "id": 2739
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "13",
          "south": "none",
          "west": "up"
        },
        "id": 2740
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "13",
          "south": "none",
          "west": "side"
        },
        "id": 2741
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "13",
          "south": "none",
          "west": "none"
        },
        "id": 2742
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "14",
          "south": "up",
          "west": "up"
        },
        "id": 2743
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "14",
          "south": "up",
          "west": "side"
        },
        "id": 2744
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "14",
          "south": "up",
          "west": "none"
        },
        "id": 2745
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "14",
          "south": "side",
          "west": "up"
        },
        "id": 2746
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "14",
          "south": "side",
          "west": "side"
        },
        "id": 2747
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "14",
          "south": "side",
          "west": "none"
        },
        "id": 2748
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "14",
          "south": "none",
          "west": "up"
        },
        "id": 2749
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "14",
          "south": "none",
          "west": "side"
        },
        "id": 2750
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "14",
          "south": "none",
          "west": "none"
        },
        "id": 2751
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "15",
          "south": "up",
          "west": "up"
        },
        "id": 2752
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "15",
          "south": "up",
          "west": "side"
        },
        "id": 2753
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "15",
          "south": "up",
          "west": "none"
        },
        "id": 2754
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "15",
          "south": "side",
          "west": "up"
        },
        "id": 2755
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "15",
          "south": "side",
          "west": "side"
        },
        "id": 2756
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "15",
          "south": "side",
          "west": "none"
        },
        "id": 2757
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "15",
          "south": "none",
          "west": "up"
        },
        "id": 2758
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "15",
          "south": "none",
          "west": "side"
        },
        "id": 2759
      },
      {
        "properties": {
          "east": "none",
          "north": "up",
          "power": "15",
          "south": "none",
          "west": "none"
        },
        "id": 2760
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "0",
          "south": "up",
          "west": "up"
        },
        "id": 2761
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "0",
          "south": "up",
          "west": "side"
        },
        "id": 2762
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "0",
          "south": "up",
          "west": "none"
        },
        "id": 2763
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "0",
          "south": "side",
          "west": "up"
        },
        "id": 2764
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "0",
          "south": "side",
          "west": "side"
        },
        "id": 2765
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "0",
          "south": "side",
          "west": "none"
        },
        "id": 2766
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "0",
          "south": "none",
          "west": "up"
        },
        "id": 2767
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "0",
          "south": "none",
          "west": "side"
        },
        "id": 2768
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "0",
          "south": "none",
          "west": "none"
        },
        "id": 2769
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "1",
          "south": "up",
          "west": "up"
        },
        "id": 2770
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "1",
          "south": "up",
          "west": "side"
        },
        "id": 2771
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "1",
          "south": "up",
          "west": "none"
        },
        "id": 2772
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "1",
          "south": "side",
          "west": "up"
        },
        "id": 2773
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "1",
          "south": "side",
          "west": "side"
        },
        "id": 2774
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "1",
          "south": "side",
          "west": "none"
        },
        "id": 2775
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "1",
          "south": "none",
          "west": "up"
        },
        "id": 2776
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "1",
          "south": "none",
          "west": "side"
        },
        "id": 2777
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "1",
          "south": "none",
          "west": "none"
        },
        "id": 2778
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "2",
          "south": "up",
          "west": "up"
        },
        "id": 2779
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "2",
          "south": "up",
          "west": "side"
        },
        "id": 2780
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "2",
          "south": "up",
          "west": "none"
        },
        "id": 2781
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "2",
          "south": "side",
          "west": "up"
        },
        "id": 2782
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "2",
          "south": "side",
          "west": "side"
        },
        "id": 2783
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "2",
          "south": "side",
          "west": "none"
        },
        "id": 2784
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "2",
          "south": "none",
          "west": "up"
        },
        "id": 2785
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "2",
          "south": "none",
          "west": "side"
        },
        "id": 2786
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "2",
          "south": "none",
          "west": "none"
        },
        "id": 2787
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "3",
          "south": "up",
          "west": "up"
        },
        "id": 2788
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "3",
          "south": "up",
          "west": "side"
        },
        "id": 2789
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "3",
          "south": "up",
          "west": "none"
        },
        "id": 2790
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "3",
          "south": "side",
          "west": "up"
        },
        "id": 2791
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "3",
          "south": "side",
          "west": "side"
        },
        "id": 2792
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "3",
          "south": "side",
          "west": "none"
        },
        "id": 2793
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "3",
          "south": "none",
          "west": "up"
        },
        "id": 2794
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "3",
          "south": "none",
          "west": "side"
        },
        "id": 2795
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "3",
          "south": "none",
          "west": "none"
        },
        "id": 2796
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "4",
          "south": "up",
          "west": "up"
        },
        "id": 2797
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "4",
          "south": "up",
          "west": "side"
        },
        "id": 2798
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "4",
          "south": "up",
          "west": "none"
        },
        "id": 2799
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "4",
          "south": "side",
          "west": "up"
        },
        "id": 2800
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "4",
          "south": "side",
          "west": "side"
        },
        "id": 2801
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "4",
          "south": "side",
          "west": "none"
        },
        "id": 2802
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "4",
          "south": "none",
          "west": "up"
        },
        "id": 2803
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "4",
          "south": "none",
          "west": "side"
        },
        "id": 2804
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "4",
          "south": "none",
          "west": "none"
        },
        "id": 2805
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "5",
          "south": "up",
          "west": "up"
        },
        "id": 2806
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "5",
          "south": "up",
          "west": "side"
        },
        "id": 2807
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "5",
          "south": "up",
          "west": "none"
        },
        "id": 2808
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "5",
          "south": "side",
          "west": "up"
        },
        "id": 2809
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "5",
          "south": "side",
          "west": "side"
        },
        "id": 2810
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "5",
          "south": "side",
          "west": "none"
        },
        "id": 2811
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "5",
          "south": "none",
          "west": "up"
        },
        "id": 2812
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "5",
          "south": "none",
          "west": "side"
        },
        "id": 2813
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "5",
          "south": "none",
          "west": "none"
        },
        "id": 2814
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "6",
          "south": "up",
          "west": "up"
        },
        "id": 2815
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "6",
          "south": "up",
          "west": "side"
        },
        "id": 2816
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "6",
          "south": "up",
          "west": "none"
        },
        "id": 2817
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "6",
          "south": "side",
          "west": "up"
        },
        "id": 2818
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "6",
          "south": "side",
          "west": "side"
        },
        "id": 2819
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "6",
          "south": "side",
          "west": "none"
        },
        "id": 2820
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "6",
          "south": "none",
          "west": "up"
        },
        "id": 2821
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "6",
          "south": "none",
          "west": "side"
        },
        "id": 2822
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "6",
          "south": "none",
          "west": "none"
        },
        "id": 2823
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "7",
          "south": "up",
          "west": "up"
        },
        "id": 2824
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "7",
          "south": "up",
          "west": "side"
        },
        "id": 2825
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "7",
          "south": "up",
          "west": "none"
        },
        "id": 2826
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "7",
          "south": "side",
          "west": "up"
        },
        "id": 2827
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "7",
          "south": "side",
          "west": "side"
        },
        "id": 2828
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "7",
          "south": "side",
          "west": "none"
        },
        "id": 2829
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "7",
          "south": "none",
          "west": "up"
        },
        "id": 2830
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "7",
          "south": "none",
          "west": "side"
        },
        "id": 2831
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "7",
          "south": "none",
          "west": "none"
        },
        "id": 2832
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "8",
          "south": "up",
          "west": "up"
        },
        "id": 2833
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "8",
          "south": "up",
          "west": "side"
        },
        "id": 2834
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "8",
          "south": "up",
          "west": "none"
        },
        "id": 2835
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "8",
          "south": "side",
          "west": "up"
        },
        "id": 2836
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "8",
          "south": "side",
          "west": "side"
        },
        "id": 2837
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "8",
          "south": "side",
          "west": "none"
        },
        "id": 2838
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "8",
          "south": "none",
          "west": "up"
        },
        "id": 2839
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "8",
          "south": "none",
          "west": "side"
        },
        "id": 2840
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "8",
          "south": "none",
          "west": "none"
        },
        "id": 2841
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "9",
          "south": "up",
          "west": "up"
        },
        "id": 2842
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "9",
          "south": "up",
          "west": "side"
        },
        "id": 2843
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "9",
          "south": "up",
          "west": "none"
        },
        "id": 2844
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "9",
          "south": "side",
          "west": "up"
        },
        "id": 2845
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "9",
          "south": "side",
          "west": "side"
        },
        "id": 2846
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "9",
          "south": "side",
          "west": "none"
        },
        "id": 2847
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "9",
          "south": "none",
          "west": "up"
        },
        "id": 2848
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "9",
          "south": "none",
          "west": "side"
        },
        "id": 2849
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "9",
          "south": "none",
          "west": "none"
        },
        "id": 2850
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "10",
          "south": "up",
          "west": "up"
        },
        "id": 2851
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "10",
          "south": "up",
          "west": "side"
        },
        "id": 2852
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "10",
          "south": "up",
          "west": "none"
        },
        "id": 2853
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "10",
          "south": "side",
          "west": "up"
        },
        "id": 2854
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "10",
          "south": "side",
          "west": "side"
        },
        "id": 2855
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "10",
          "south": "side",
          "west": "none"
        },
        "id": 2856
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "10",
          "south": "none",
          "west": "up"
        },
        "id": 2857
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "10",
          "south": "none",
          "west": "side"
        },
        "id": 2858
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "10",
          "south": "none",
          "west": "none"
        },
        "id": 2859
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "11",
          "south": "up",
          "west": "up"
        },
        "id": 2860
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "11",
          "south": "up",
          "west": "side"
        },
        "id": 2861
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "11",
          "south": "up",
          "west": "none"
        },
        "id": 2862
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "11",
          "south": "side",
          "west": "up"
        },
        "id": 2863
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "11",
          "south": "side",
          "west": "side"
        },
        "id": 2864
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "11",
          "south": "side",
          "west": "none"
        },
        "id": 2865
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "11",
          "south": "none",
          "west": "up"
        },
        "id": 2866
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "11",
          "south": "none",
          "west": "side"
        },
        "id": 2867
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "11",
          "south": "none",
          "west": "none"
        },
        "id": 2868
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "12",
          "south": "up",
          "west": "up"
        },
        "id": 2869
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "12",
          "south": "up",
          "west": "side"
        },
        "id": 2870
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "12",
          "south": "up",
          "west": "none"
        },
        "id": 2871
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "12",
          "south": "side",
          "west": "up"
        },
        "id": 2872
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "12",
          "south": "side",
          "west": "side"
        },
        "id": 2873
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "12",
          "south": "side",
          "west": "none"
        },
        "id": 2874
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "12",
          "south": "none",
          "west": "up"
        },
        "id": 2875
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "12",
          "south": "none",
          "west": "side"
        },
        "id": 2876
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "12",
          "south": "none",
          "west": "none"
        },
        "id": 2877
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "13",
          "south": "up",
          "west": "up"
        },
        "id": 2878
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "13",
          "south": "up",
          "west": "side"
        },
        "id": 2879
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "13",
          "south": "up",
          "west": "none"
        },
        "id": 2880
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "13",
          "south": "side",
          "west": "up"
        },
        "id": 2881
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "13",
          "south": "side",
          "west": "side"
        },
        "id": 2882
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "13",
          "south": "side",
          "west": "none"
        },
        "id": 2883
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "13",
          "south": "none",
          "west": "up"
        },
        "id": 2884
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "13",
          "south": "none",
          "west": "side"
        },
        "id": 2885
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "13",
          "south": "none",
          "west": "none"
        },
        "id": 2886
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "14",
          "south": "up",
          "west": "up"
        },
        "id": 2887
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "14",
          "south": "up",
          "west": "side"
        },
        "id": 2888
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "14",
          "south": "up",
          "west": "none"
        },
        "id": 2889
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "14",
          "south": "side",
          "west": "up"
        },
        "id": 2890
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "14",
          "south": "side",
          "west": "side"
        },
        "id": 2891
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "14",
          "south": "side",
          "west": "none"
        },
        "id": 2892
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "14",
          "south": "none",
          "west": "up"
        },
        "id": 2893
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "14",
          "south": "none",
          "west": "side"
        },
        "id": 2894
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "14",
          "south": "none",
          "west": "none"
        },
        "id": 2895
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "15",
          "south": "up",
          "west": "up"
        },
        "id": 2896
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "15",
          "south": "up",
          "west": "side"
        },
        "id": 2897
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "15",
          "south": "up",
          "west": "none"
        },
        "id": 2898
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "15",
          "south": "side",
          "west": "up"
        },
        "id": 2899
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "15",
          "south": "side",
          "west": "side"
        },
        "id": 2900
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "15",
          "south": "side",
          "west": "none"
        },
        "id": 2901
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "15",
          "south": "none",
          "west": "up"
        },
        "id": 2902
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "15",
          "south": "none",
          "west": "side"
        },
        "id": 2903
      },
      {
        "properties": {
          "east": "none",
          "north": "side",
          "power": "15",
          "south": "none",
          "west": "none"
        },
        "id": 2904
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "0",
          "south": "up",
          "west": "up"
        },
        "id": 2905
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "0",
          "south": "up",
          "west": "side"
        },
        "id": 2906
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "0",
          "south": "up",
          "west": "none"
        },
        "id": 2907
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "0",
          "south": "side",
          "west": "up"
        },
        "id": 2908
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "0",
          "south": "side",
          "west": "side"
        },
        "id": 2909
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "0",
          "south": "side",
          "west": "none"
        },
        "id": 2910
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "0",
          "south": "none",
          "west": "up"
        },
        "id": 2911
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "0",
          "south": "none",
          "west": "side"
        },
        "id": 2912
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "0",
          "south": "none",
          "west": "none"
        },
        "id": 2913,
        "default": true
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "1",
          "south": "up",
          "west": "up"
        },
        "id": 2914
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "1",
          "south": "up",
          "west": "side"
        },
        "id": 2915
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "1",
          "south": "up",
          "west": "none"
        },
        "id": 2916
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "1",
          "south": "side",
          "west": "up"
        },
        "id": 2917
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "1",
          "south": "side",
          "west": "side"
        },
        "id": 2918
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "1",
          "south": "side",
          "west": "none"
        },
        "id": 2919
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "1",
          "south": "none",
          "west": "up"
        },
        "id": 2920
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "1",
          "south": "none",
          "west": "side"
        },
        "id": 2921
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "1",
          "south": "none",
          "west": "none"
        },
        "id": 2922
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "2",
          "south": "up",
          "west": "up"
        },
        "id": 2923
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "2",
          "south": "up",
          "west": "side"
        },
        "id": 2924
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "2",
          "south": "up",
          "west": "none"
        },
        "id": 2925
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "2",
          "south": "side",
          "west": "up"
        },
        "id": 2926
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "2",
          "south": "side",
          "west": "side"
        },
        "id": 2927
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "2",
          "south": "side",
          "west": "none"
        },
        "id": 2928
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "2",
          "south": "none",
          "west": "up"
        },
        "id": 2929
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "2",
          "south": "none",
          "west": "side"
        },
        "id": 2930
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "2",
          "south": "none",
          "west": "none"
        },
        "id": 2931
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "3",
          "south": "up",
          "west": "up"
        },
        "id": 2932
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "3",
          "south": "up",
          "west": "side"
        },
        "id": 2933
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "3",
          "south": "up",
          "west": "none"
        },
        "id": 2934
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "3",
          "south": "side",
          "west": "up"
        },
        "id": 2935
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "3",
          "south": "side",
          "west": "side"
        },
        "id": 2936
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "3",
          "south": "side",
          "west": "none"
        },
        "id": 2937
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "3",
          "south": "none",
          "west": "up"
        },
        "id": 2938
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "3",
          "south": "none",
          "west": "side"
        },
        "id": 2939
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "3",
          "south": "none",
          "west": "none"
        },
        "id": 2940
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "4",
          "south": "up",
          "west": "up"
        },
        "id": 2941
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "4",
          "south": "up",
          "west": "side"
        },
        "id": 2942
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "4",
          "south": "up",
          "west": "none"
        },
        "id": 2943
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "4",
          "south": "side",
          "west": "up"
        },
        "id": 2944
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "4",
          "south": "side",
          "west": "side"
        },
        "id": 2945
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "4",
          "south": "side",
          "west": "none"
        },
        "id": 2946
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "4",
          "south": "none",
          "west": "up"
        },
        "id": 2947
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "4",
          "south": "none",
          "west": "side"
        },
        "id": 2948
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "4",
          "south": "none",
          "west": "none"
        },
        "id": 2949
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "5",
          "south": "up",
          "west": "up"
        },
        "id": 2950
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "5",
          "south": "up",
          "west": "side"
        },
        "id": 2951
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "5",
          "south": "up",
          "west": "none"
        },
        "id": 2952
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "5",
          "south": "side",
          "west": "up"
        },
        "id": 2953
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "5",
          "south": "side",
          "west": "side"
        },
        "id": 2954
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "5",
          "south": "side",
          "west": "none"
        },
        "id": 2955
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "5",
          "south": "none",
          "west": "up"
        },
        "id": 2956
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "5",
          "south": "none",
          "west": "side"
        },
        "id": 2957
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "5",
          "south": "none",
          "west": "none"
        },
        "id": 2958
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "6",
          "south": "up",
          "west": "up"
        },
        "id": 2959
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "6",
          "south": "up",
          "west": "side"
        },
        "id": 2960
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "6",
          "south": "up",
          "west": "none"
        },
        "id": 2961
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "6",
          "south": "side",
          "west": "up"
        },
        "id": 2962
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "6",
          "south": "side",
          "west": "side"
        },
        "id": 2963
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "6",
          "south": "side",
          "west": "none"
        },
        "id": 2964
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "6",
          "south": "none",
          "west": "up"
        },
        "id": 2965
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "6",
          "south": "none",
          "west": "side"
        },
        "id": 2966
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "6",
          "south": "none",
          "west": "none"
        },
        "id": 2967
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "7",
          "south": "up",
          "west": "up"
        },
        "id": 2968
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "7",
          "south": "up",
          "west": "side"
        },
        "id": 2969
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "7",
          "south": "up",
          "west": "none"
        },
        "id": 2970
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "7",
          "south": "side",
          "west": "up"
        },
        "id": 2971
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "7",
          "south": "side",
          "west": "side"
        },
        "id": 2972
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "7",
          "south": "side",
          "west": "none"
        },
        "id": 2973
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "7",
          "south": "none",
          "west": "up"
        },
        "id": 2974
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "7",
          "south": "none",
          "west": "side"
        },
        "id": 2975
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "7",
          "south": "none",
          "west": "none"
        },
        "id": 2976
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "8",
          "south": "up",
          "west": "up"
        },
        "id": 2977
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "8",
          "south": "up",
          "west": "side"
        },
        "id": 2978
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "8",
          "south": "up",
          "west": "none"
        },
        "id": 2979
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "8",
          "south": "side",
          "west": "up"
        },
        "id": 2980
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "8",
          "south": "side",
          "west": "side"
        },
        "id": 2981
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "8",
          "south": "side",
          "west": "none"
        },
        "id": 2982
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "8",
          "south": "none",
          "west": "up"
        },
        "id": 2983
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "8",
          "south": "none",
          "west": "side"
        },
        "id": 2984
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "8",
          "south": "none",
          "west": "none"
        },
        "id": 2985
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "9",
          "south": "up",
          "west": "up"
        },
        "id": 2986
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "9",
          "south": "up",
          "west": "side"
        },
        "id": 2987
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "9",
          "south": "up",
          "west": "none"
        },
        "id": 2988
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "9",
          "south": "side",
          "west": "up"
        },
        "id": 2989
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "9",
          "south": "side",
          "west": "side"
        },
        "id": 2990
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "9",
          "south": "side",
          "west": "none"
        },
        "id": 2991
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "9",
          "south": "none",
          "west": "up"
        },
        "id": 2992
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "9",
          "south": "none",
          "west": "side"
        },
        "id": 2993
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "9",
          "south": "none",
          "west": "none"
        },
        "id": 2994
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "10",
          "south": "up",
          "west": "up"
        },
        "id": 2995
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "10",
          "south": "up",
          "west": "side"
        },
        "id": 2996
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "10",
          "south": "up",
          "west": "none"
        },
        "id": 2997
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "10",
          "south": "side",
          "west": "up"
        },
        "id": 2998
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "10",
          "south": "side",
          "west": "side"
        },
        "id": 2999
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "10",
          "south": "side",
          "west": "none"
        },
        "id": 3000
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "10",
          "south": "none",
          "west": "up"
        },
        "id": 3001
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "10",
          "south": "none",
          "west": "side"
        },
        "id": 3002
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "10",
          "south": "none",
          "west": "none"
        },
        "id": 3003
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "11",
          "south": "up",
          "west": "up"
        },
        "id": 3004
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "11",
          "south": "up",
          "west": "side"
        },
        "id": 3005
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "11",
          "south": "up",
          "west": "none"
        },
        "id": 3006
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "11",
          "south": "side",
          "west": "up"
        },
        "id": 3007
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "11",
          "south": "side",
          "west": "side"
        },
        "id": 3008
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "11",
          "south": "side",
          "west": "none"
        },
        "id": 3009
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "11",
          "south": "none",
          "west": "up"
        },
        "id": 3010
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "11",
          "south": "none",
          "west": "side"
        },
        "id": 3011
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "11",
          "south": "none",
          "west": "none"
        },
        "id": 3012
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "12",
          "south": "up",
          "west": "up"
        },
        "id": 3013
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "12",
          "south": "up",
          "west": "side"
        },
        "id": 3014
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "12",
          "south": "up",
          "west": "none"
        },
        "id": 3015
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "12",
          "south": "side",
          "west": "up"
        },
        "id": 3016
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "12",
          "south": "side",
          "west": "side"
        },
        "id": 3017
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "12",
          "south": "side",
          "west": "none"
        },
        "id": 3018
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "12",
          "south": "none",
          "west": "up"
        },
        "id": 3019
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "12",
          "south": "none",
          "west": "side"
        },
        "id": 3020
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "12",
          "south": "none",
          "west": "none"
        },
        "id": 3021
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "13",
          "south": "up",
          "west": "up"
        },
        "id": 3022
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "13",
          "south": "up",
          "west": "side"
        },
        "id": 3023
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "13",
          "south": "up",
          "west": "none"
        },
        "id": 3024
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "13",
          "south": "side",
          "west": "up"
        },
        "id": 3025
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "13",
          "south": "side",
          "west": "side"
        },
        "id": 3026
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "13",
          "south": "side",
          "west": "none"
        },
        "id": 3027
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "13",
          "south": "none",
          "west": "up"
        },
        "id": 3028
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "13",
          "south": "none",
          "west": "side"
        },
        "id": 3029
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "13",
          "south": "none",
          "west": "none"
        },
        "id": 3030
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "14",
          "south": "up",
          "west": "up"
        },
        "id": 3031
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "14",
          "south": "up",
          "west": "side"
        },
        "id": 3032
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "14",
          "south": "up",
          "west": "none"
        },
        "id": 3033
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "14",
          "south": "side",
          "west": "up"
        },
        "id": 3034
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "14",
          "south": "side",
          "west": "side"
        },
        "id": 3035
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "14",
          "south": "side",
          "west": "none"
        },
        "id": 3036
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "14",
          "south": "none",
          "west": "up"
        },
        "id": 3037
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "14",
          "south": "none",
          "west": "side"
        },
        "id": 3038
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "14",
          "south": "none",
          "west": "none"
        },
        "id": 3039
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "15",
          "south": "up",
          "west": "up"
        },
        "id": 3040
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "15",
          "south": "up",
          "west": "side"
        },
        "id": 3041
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "15",
          "south": "up",
          "west": "none"
        },
        "id": 3042
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "15",
          "south": "side",
          "west": "up"
        },
        "id": 3043
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "15",
          "south": "side",
          "west": "side"
        },
        "id": 3044
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "15",
          "south": "side",
          "west": "none"
        },
        "id": 3045
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "15",
          "south": "none",
          "west": "up"
        },
        "id": 3046
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "15",
          "south": "none",
          "west": "side"
        },
        "id": 3047
      },
      {
        "properties": {
          "east": "none",
          "north": "none",
          "power": "15",
          "south": "none",
          "west": "none"
        },
        "id": 3048
      }
    ]
  },
  "minecraft:diamond_ore": {
    "states": [
      {
        "id": 3049,
        "default": true
      }
    ]
  },
  "minecraft:diamond_block": {
    "states": [
      {
        "id": 3050,
        "default": true
      }
    ]
  },
  "minecraft:crafting_table": {
    "states": [
      {
        "id": 3051,
        "default": true
      }
    ]
  },
  "minecraft:wheat": {
    "properties": {
      "age": [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7"
      ]
    },
    "states": [
      {
        "properties": {
          "age": "0"
        },
        "id": 3052,
        "default": true
      },
      {
        "properties": {
          "age": "1"
        },
        "id": 3053
      },
      {
        "properties": {
          "age": "2"
        },
        "id": 3054
      },
      {
        "properties": {
          "age": "3"
        },
        "id": 3055
      },
      {
        "properties": {
          "age": "4"
        },
        "id": 3056
      },
      {
        "properties": {
          "age": "5"
        },
        "id": 3057
      },
      {
        "properties": {
          "age": "6"
        },
        "id": 3058
      },
      {
        "properties": {
          "age": "7"
        },
        "id": 3059
      }
    ]
  },
  "minecraft:farmland": {
    "properties": {
      "moisture": [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7"
      ]
    },
    "states": [
      {
        "properties": {
          "moisture": "0"
        },
        "id": 3060,
        "default": true
      },
      {
        "properties": {
          "moisture": "1"
        },
        "id": 3061
      },
      {
        "properties": {
          "moisture": "2"
        },
        "id": 3062
      },
      {
        "properties": {
          "moisture": "3"
        },
        "id": 3063
      },
      {
        "properties": {
          "moisture": "4"
        },
        "id": 3064
      },
      {
        "properties": {
          "moisture": "5"
        },
        "id": 3065
      },
      {
        "properties": {
          "moisture": "6"
        },
        "id": 3066
      },
      {
        "properties": {
          "moisture": "7"
        },
        "id": 3067
      }
    ]
  },
  "minecraft:furnace": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "lit": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "lit": "true"
        },
        "id": 3068
      },
      {
        "properties": {
          "facing": "north",
          "lit": "false"
        },
        "id": 3069,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "lit": "true"
        },
        "id": 3070
      },
      {
        "properties": {
          "facing": "south",
          "lit": "false"
        },
        "id": 3071
      },
      {
        "properties": {
          "facing": "west",
          "lit": "true"
        },
        "id": 3072
      },
      {
        "properties": {
          "facing": "west",
          "lit": "false"
        },
        "id": 3073
      },
      {
        "properties": {
          "facing": "east",
          "lit": "true"
        },
        "id": 3074
      },
      {
        "properties": {
          "facing": "east",
          "lit": "false"
        },
        "id": 3075
      }
    ]
  },
  "minecraft:sign": {
    "properties": {
      "rotation": [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "10",
        "11",
        "12",
        "13",
        "14",
        "15"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "rotation": "0",
          "waterlogged": "true"
        },
        "id": 3076
      },
      {
        "properties": {
          "rotation": "0",
          "waterlogged": "false"
        },
        "id": 3077,
        "default": true
      },
      {
        "properties": {
          "rotation": "1",
          "waterlogged": "true"
        },
        "id": 3078
      },
      {
        "properties": {
          "rotation": "1",
          "waterlogged": "false"
        },
        "id": 3079
      },
      {
        "properties": {
          "rotation": "2",
          "waterlogged": "true"
        },
        "id": 3080
      },
      {
        "properties": {
          "rotation": "2",
          "waterlogged": "false"
        },
        "id": 3081
      },
      {
        "properties": {
          "rotation": "3",
          "waterlogged": "true"
        },
        "id": 3082
      },
      {
        "properties": {
          "rotation": "3",
          "waterlogged": "false"
        },
        "id": 3083
      },
      {
        "properties": {
          "rotation": "4",
          "waterlogged": "true"
        },
        "id": 3084
      },
      {
        "properties": {
          "rotation": "4",
          "waterlogged": "false"
        },
        "id": 3085
      },
      {
        "properties": {
          "rotation": "5",
          "waterlogged": "true"
        },
        "id": 3086
      },
      {
        "properties": {
          "rotation": "5",
          "waterlogged": "false"
        },
        "id": 3087
      },
      {
        "properties": {
          "rotation": "6",
          "waterlogged": "true"
        },
        "id": 3088
      },
      {
        "properties": {
          "rotation": "6",
          "waterlogged": "false"
        },
        "id": 3089
      },
      {
        "properties": {
          "rotation": "7",
          "waterlogged": "true"
        },
        "id": 3090
      },
      {
        "properties": {
          "rotation": "7",
          "waterlogged": "false"
        },
        "id": 3091
      },
      {
        "properties": {
          "rotation": "8",
          "waterlogged": "true"
        },
        "id": 3092
      },
      {
        "properties": {
          "rotation": "8",
          "waterlogged": "false"
        },
        "id": 3093
      },
      {
        "properties": {
          "rotation": "9",
          "waterlogged": "true"
        },
        "id": 3094
      },
      {
        "properties": {
          "rotation": "9",
          "waterlogged": "false"
        },
        "id": 3095
      },
      {
        "properties": {
          "rotation": "10",
          "waterlogged": "true"
        },
        "id": 3096
      },
      {
        "properties": {
          "rotation": "10",
          "waterlogged": "false"
        },
        "id": 3097
      },
      {
        "properties": {
          "rotation": "11",
          "waterlogged": "true"
        },
        "id": 3098
      },
      {
        "properties": {
          "rotation": "11",
          "waterlogged": "false"
        },
        "id": 3099
      },
      {
        "properties": {
          "rotation": "12",
          "waterlogged": "true"
        },
        "id": 3100
      },
      {
        "properties": {
          "rotation": "12",
          "waterlogged": "false"
        },
        "id": 3101
      },
      {
        "properties": {
          "rotation": "13",
          "waterlogged": "true"
        },
        "id": 3102
      },
      {
        "properties": {
          "rotation": "13",
          "waterlogged": "false"
        },
        "id": 3103
      },
      {
        "properties": {
          "rotation": "14",
          "waterlogged": "true"
        },
        "id": 3104
      },
      {
        "properties": {
          "rotation": "14",
          "waterlogged": "false"
        },
        "id": 3105
      },
      {
        "properties": {
          "rotation": "15",
          "waterlogged": "true"
        },
        "id": 3106
      },
      {
        "properties": {
          "rotation": "15",
          "waterlogged": "false"
        },
        "id": 3107
      }
    ]
  },
  "minecraft:oak_door": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "upper",
        "lower"
      ],
      "hinge": [
        "left",
        "right"
      ],
      "open": [
        "true",
        "false"
      ],
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3108
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3109
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3110
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3111
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3112
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3113
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3114
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3115
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3116
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3117
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3118
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3119,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3120
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3121
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3122
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3123
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3124
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3125
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3126
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3127
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3128
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3129
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3130
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3131
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3132
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3133
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3134
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3135
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3136
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3137
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3138
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3139
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3140
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3141
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3142
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3143
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3144
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3145
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3146
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3147
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3148
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3149
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3150
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3151
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3152
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3153
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3154
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3155
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3156
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3157
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3158
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3159
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3160
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3161
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3162
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3163
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3164
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3165
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3166
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3167
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3168
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3169
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3170
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3171
      }
    ]
  },
  "minecraft:ladder": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "waterlogged": "true"
        },
        "id": 3172
      },
      {
        "properties": {
          "facing": "north",
          "waterlogged": "false"
        },
        "id": 3173,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "waterlogged": "true"
        },
        "id": 3174
      },
      {
        "properties": {
          "facing": "south",
          "waterlogged": "false"
        },
        "id": 3175
      },
      {
        "properties": {
          "facing": "west",
          "waterlogged": "true"
        },
        "id": 3176
      },
      {
        "properties": {
          "facing": "west",
          "waterlogged": "false"
        },
        "id": 3177
      },
      {
        "properties": {
          "facing": "east",
          "waterlogged": "true"
        },
        "id": 3178
      },
      {
        "properties": {
          "facing": "east",
          "waterlogged": "false"
        },
        "id": 3179
      }
    ]
  },
  "minecraft:rail": {
    "properties": {
      "shape": [
        "north_south",
        "east_west",
        "ascending_east",
        "ascending_west",
        "ascending_north",
        "ascending_south",
        "south_east",
        "south_west",
        "north_west",
        "north_east"
      ]
    },
    "states": [
      {
        "properties": {
          "shape": "north_south"
        },
        "id": 3180,
        "default": true
      },
      {
        "properties": {
          "shape": "east_west"
        },
        "id": 3181
      },
      {
        "properties": {
          "shape": "ascending_east"
        },
        "id": 3182
      },
      {
        "properties": {
          "shape": "ascending_west"
        },
        "id": 3183
      },
      {
        "properties": {
          "shape": "ascending_north"
        },
        "id": 3184
      },
      {
        "properties": {
          "shape": "ascending_south"
        },
        "id": 3185
      },
      {
        "properties": {
          "shape": "south_east"
        },
        "id": 3186
      },
      {
        "properties": {
          "shape": "south_west"
        },
        "id": 3187
      },
      {
        "properties": {
          "shape": "north_west"
        },
        "id": 3188
      },
      {
        "properties": {
          "shape": "north_east"
        },
        "id": 3189
      }
    ]
  },
  "minecraft:cobblestone_stairs": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "shape": [
        "straight",
        "inner_left",
        "inner_right",
        "outer_left",
        "outer_right"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 3190
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 3191
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 3192
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 3193
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 3194
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 3195
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 3196
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 3197
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 3198
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 3199
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 3200
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 3201,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 3202
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 3203
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 3204
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 3205
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 3206
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 3207
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 3208
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 3209
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 3210
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 3211
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 3212
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 3213
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 3214
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 3215
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 3216
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 3217
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 3218
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 3219
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 3220
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 3221
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 3222
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 3223
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 3224
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 3225
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 3226
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 3227
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 3228
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 3229
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 3230
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 3231
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 3232
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 3233
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 3234
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 3235
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 3236
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 3237
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 3238
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 3239
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 3240
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 3241
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 3242
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 3243
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 3244
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 3245
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 3246
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 3247
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 3248
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 3249
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 3250
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 3251
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 3252
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 3253
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 3254
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 3255
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 3256
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 3257
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 3258
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 3259
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 3260
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 3261
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 3262
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 3263
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 3264
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 3265
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 3266
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 3267
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 3268
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 3269
      }
    ]
  },
  "minecraft:wall_sign": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "waterlogged": "true"
        },
        "id": 3270
      },
      {
        "properties": {
          "facing": "north",
          "waterlogged": "false"
        },
        "id": 3271,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "waterlogged": "true"
        },
        "id": 3272
      },
      {
        "properties": {
          "facing": "south",
          "waterlogged": "false"
        },
        "id": 3273
      },
      {
        "properties": {
          "facing": "west",
          "waterlogged": "true"
        },
        "id": 3274
      },
      {
        "properties": {
          "facing": "west",
          "waterlogged": "false"
        },
        "id": 3275
      },
      {
        "properties": {
          "facing": "east",
          "waterlogged": "true"
        },
        "id": 3276
      },
      {
        "properties": {
          "facing": "east",
          "waterlogged": "false"
        },
        "id": 3277
      }
    ]
  },
  "minecraft:lever": {
    "properties": {
      "face": [
        "floor",
        "wall",
        "ceiling"
      ],
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "face": "floor",
          "facing": "north",
          "powered": "true"
        },
        "id": 3278
      },
      {
        "properties": {
          "face": "floor",
          "facing": "north",
          "powered": "false"
        },
        "id": 3279
      },
      {
        "properties": {
          "face": "floor",
          "facing": "south",
          "powered": "true"
        },
        "id": 3280
      },
      {
        "properties": {
          "face": "floor",
          "facing": "south",
          "powered": "false"
        },
        "id": 3281
      },
      {
        "properties": {
          "face": "floor",
          "facing": "west",
          "powered": "true"
        },
        "id": 3282
      },
      {
        "properties": {
          "face": "floor",
          "facing": "west",
          "powered": "false"
        },
        "id": 3283
      },
      {
        "properties": {
          "face": "floor",
          "facing": "east",
          "powered": "true"
        },
        "id": 3284
      },
      {
        "properties": {
          "face": "floor",
          "facing": "east",
          "powered": "false"
        },
        "id": 3285
      },
      {
        "properties": {
          "face": "wall",
          "facing": "north",
          "powered": "true"
        },
        "id": 3286
      },
      {
        "properties": {
          "face": "wall",
          "facing": "north",
          "powered": "false"
        },
        "id": 3287,
        "default": true
      },
      {
        "properties": {
          "face": "wall",
          "facing": "south",
          "powered": "true"
        },
        "id": 3288
      },
      {
        "properties": {
          "face": "wall",
          "facing": "south",
          "powered": "false"
        },
        "id": 3289
      },
      {
        "properties": {
          "face": "wall",
          "facing": "west",
          "powered": "true"
        },
        "id": 3290
      },
      {
        "properties": {
          "face": "wall",
          "facing": "west",
          "powered": "false"
        },
        "id": 3291
      },
      {
        "properties": {
          "face": "wall",
          "facing": "east",
          "powered": "true"
        },
        "id": 3292
      },
      {
        "properties": {
          "face": "wall",
          "facing": "east",
          "powered": "false"
        },
        "id": 3293
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "north",
          "powered": "true"
        },
        "id": 3294
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "north",
          "powered": "false"
        },
        "id": 3295
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "south",
          "powered": "true"
        },
        "id": 3296
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "south",
          "powered": "false"
        },
        "id": 3297
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "west",
          "powered": "true"
        },
        "id": 3298
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "west",
          "powered": "false"
        },
        "id": 3299
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "east",
          "powered": "true"
        },
        "id": 3300
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "east",
          "powered": "false"
        },
        "id": 3301
      }
    ]
  },
  "minecraft:stone_pressure_plate": {
    "properties": {
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "powered": "true"
        },
        "id": 3302
      },
      {
        "properties": {
          "powered": "false"
        },
        "id": 3303,
        "default": true
      }
    ]
  },
  "minecraft:iron_door": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "upper",
        "lower"
      ],
      "hinge": [
        "left",
        "right"
      ],
      "open": [
        "true",
        "false"
      ],
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3304
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3305
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3306
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3307
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3308
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3309
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3310
      },
      {
        "properties": {
          "facing": "north",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3311
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3312
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3313
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3314
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3315,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3316
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3317
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3318
      },
      {
        "properties": {
          "facing": "north",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3319
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3320
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3321
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3322
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3323
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3324
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3325
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3326
      },
      {
        "properties": {
          "facing": "south",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3327
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3328
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3329
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3330
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3331
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3332
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3333
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3334
      },
      {
        "properties": {
          "facing": "south",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3335
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3336
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3337
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3338
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3339
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3340
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3341
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3342
      },
      {
        "properties": {
          "facing": "west",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3343
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3344
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3345
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3346
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3347
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3348
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3349
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3350
      },
      {
        "properties": {
          "facing": "west",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3351
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3352
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3353
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3354
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3355
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3356
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3357
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3358
      },
      {
        "properties": {
          "facing": "east",
          "half": "upper",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3359
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "true"
        },
        "id": 3360
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "left",
          "open": "true",
          "powered": "false"
        },
        "id": 3361
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "true"
        },
        "id": 3362
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "left",
          "open": "false",
          "powered": "false"
        },
        "id": 3363
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "true"
        },
        "id": 3364
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "right",
          "open": "true",
          "powered": "false"
        },
        "id": 3365
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "true"
        },
        "id": 3366
      },
      {
        "properties": {
          "facing": "east",
          "half": "lower",
          "hinge": "right",
          "open": "false",
          "powered": "false"
        },
        "id": 3367
      }
    ]
  },
  "minecraft:oak_pressure_plate": {
    "properties": {
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "powered": "true"
        },
        "id": 3368
      },
      {
        "properties": {
          "powered": "false"
        },
        "id": 3369,
        "default": true
      }
    ]
  },
  "minecraft:spruce_pressure_plate": {
    "properties": {
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "powered": "true"
        },
        "id": 3370
      },
      {
        "properties": {
          "powered": "false"
        },
        "id": 3371,
        "default": true
      }
    ]
  },
  "minecraft:birch_pressure_plate": {
    "properties": {
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "powered": "true"
        },
        "id": 3372
      },
      {
        "properties": {
          "powered": "false"
        },
        "id": 3373,
        "default": true
      }
    ]
  },
  "minecraft:jungle_pressure_plate": {
    "properties": {
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "powered": "true"
        },
        "id": 3374
      },
      {
        "properties": {
          "powered": "false"
        },
        "id": 3375,
        "default": true
      }
    ]
  },
  "minecraft:acacia_pressure_plate": {
    "properties": {
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "powered": "true"
        },
        "id": 3376
      },
      {
        "properties": {
          "powered": "false"
        },
        "id": 3377,
        "default": true
      }
    ]
  },
  "minecraft:dark_oak_pressure_plate": {
    "properties": {
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "powered": "true"
        },
        "id": 3378
      },
      {
        "properties": {
          "powered": "false"
        },
        "id": 3379,
        "default": true
      }
    ]
  },
  "minecraft:redstone_ore": {
    "properties": {
      "lit": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "lit": "true"
        },
        "id": 3380
      },
      {
        "properties": {
          "lit": "false"
        },
        "id": 3381,
        "default": true
      }
    ]
  },
  "minecraft:redstone_torch": {
    "properties": {
      "lit": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "lit": "true"
        },
        "id": 3382,
        "default": true
      },
      {
        "properties": {
          "lit": "false"
        },
        "id": 3383
      }
    ]
  },
  "minecraft:redstone_wall_torch": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "lit": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "lit": "true"
        },
        "id": 3384,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "lit": "false"
        },
        "id": 3385
      },
      {
        "properties": {
          "facing": "south",
          "lit": "true"
        },
        "id": 3386
      },
      {
        "properties": {
          "facing": "south",
          "lit": "false"
        },
        "id": 3387
      },
      {
        "properties": {
          "facing": "west",
          "lit": "true"
        },
        "id": 3388
      },
      {
        "properties": {
          "facing": "west",
          "lit": "false"
        },
        "id": 3389
      },
      {
        "properties": {
          "facing": "east",
          "lit": "true"
        },
        "id": 3390
      },
      {
        "properties": {
          "facing": "east",
          "lit": "false"
        },
        "id": 3391
      }
    ]
  },
  "minecraft:stone_button": {
    "properties": {
      "face": [
        "floor",
        "wall",
        "ceiling"
      ],
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "face": "floor",
          "facing": "north",
          "powered": "true"
        },
        "id": 3392
      },
      {
        "properties": {
          "face": "floor",
          "facing": "north",
          "powered": "false"
        },
        "id": 3393
      },
      {
        "properties": {
          "face": "floor",
          "facing": "south",
          "powered": "true"
        },
        "id": 3394
      },
      {
        "properties": {
          "face": "floor",
          "facing": "south",
          "powered": "false"
        },
        "id": 3395
      },
      {
        "properties": {
          "face": "floor",
          "facing": "west",
          "powered": "true"
        },
        "id": 3396
      },
      {
        "properties": {
          "face": "floor",
          "facing": "west",
          "powered": "false"
        },
        "id": 3397
      },
      {
        "properties": {
          "face": "floor",
          "facing": "east",
          "powered": "true"
        },
        "id": 3398
      },
      {
        "properties": {
          "face": "floor",
          "facing": "east",
          "powered": "false"
        },
        "id": 3399
      },
      {
        "properties": {
          "face": "wall",
          "facing": "north",
          "powered": "true"
        },
        "id": 3400
      },
      {
        "properties": {
          "face": "wall",
          "facing": "north",
          "powered": "false"
        },
        "id": 3401,
        "default": true
      },
      {
        "properties": {
          "face": "wall",
          "facing": "south",
          "powered": "true"
        },
        "id": 3402
      },
      {
        "properties": {
          "face": "wall",
          "facing": "south",
          "powered": "false"
        },
        "id": 3403
      },
      {
        "properties": {
          "face": "wall",
          "facing": "west",
          "powered": "true"
        },
        "id": 3404
      },
      {
        "properties": {
          "face": "wall",
          "facing": "west",
          "powered": "false"
        },
        "id": 3405
      },
      {
        "properties": {
          "face": "wall",
          "facing": "east",
          "powered": "true"
        },
        "id": 3406
      },
      {
        "properties": {
          "face": "wall",
          "facing": "east",
          "powered": "false"
        },
        "id": 3407
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "north",
          "powered": "true"
        },
        "id": 3408
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "north",
          "powered": "false"
        },
        "id": 3409
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "south",
          "powered": "true"
        },
        "id": 3410
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "south",
          "powered": "false"
        },
        "id": 3411
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "west",
          "powered": "true"
        },
        "id": 3412
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "west",
          "powered": "false"
        },
        "id": 3413
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "east",
          "powered": "true"
        },
        "id": 3414
      },
      {
        "properties": {
          "face": "ceiling",
          "facing": "east",
          "powered": "false"
        },
        "id": 3415
      }
    ]
  },
  "minecraft:snow": {
    "properties": {
      "layers": [
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8"
      ]
    },
    "states": [
      {
        "properties": {
          "layers": "1"
        },
        "id": 3416,
        "default": true
      },
      {
        "properties": {
          "layers": "2"
        },
        "id": 3417
      },
      {
        "properties": {
          "layers": "3"
        },
        "id": 3418
      },
      {
        "properties": {
          "layers": "4"
        },
        "id": 3419
      },
      {
        "properties": {
          "layers": "5"
        },
        "id": 3420
      },
      {
        "properties": {
          "layers": "6"
        },
        "id": 3421
      },
      {
        "properties": {
          "layers": "7"
        },
        "id": 3422
      },
      {
        "properties": {
          "layers": "8"
        },
        "id": 3423
      }
    ]
  },
  "minecraft:ice": {
    "states": [
      {
        "id": 3424,
        "default": true
      }
    ]
  },
  "minecraft:snow_block": {
    "states": [
      {
        "id": 3425,
        "default": true
      }
    ]
  },
  "minecraft:cactus": {
    "properties": {
      "age": [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "10",
        "11",
        "12",
        "13",
        "14",
        "15"
      ]
    },
    "states": [
      {
        "properties": {
          "age": "0"
        },
        "id": 3426,
        "default": true
      },
      {
        "properties": {
          "age": "1"
        },
        "id": 3427
      },
      {
        "properties": {
          "age": "2"
        },
        "id": 3428
      },
      {
        "properties": {
          "age": "3"
        },
        "id": 3429
      },
      {
        "properties": {
          "age": "4"
        },
        "id": 3430
      },
      {
        "properties": {
          "age": "5"
        },
        "id": 3431
      },
      {
        "properties": {
          "age": "6"
        },
        "id": 3432
      },
      {
        "properties": {
          "age": "7"
        },
        "id": 3433
      },
      {
        "properties": {
          "age": "8"
        },
        "id": 3434
      },
      {
        "properties": {
          "age": "9"
        },
        "id": 3435
      },
      {
        "properties": {
          "age": "10"
        },
        "id": 3436
      },
      {
        "properties": {
          "age": "11"
        },
        "id": 3437
      },
      {
        "properties": {
          "age": "12"
        },
        "id": 3438
      },
      {
        "properties": {
          "age": "13"
        },
        "id": 3439
      },
      {
        "properties": {
          "age": "14"
        },
        "id": 3440
      },
      {
        "properties": {
          "age": "15"
        },
        "id": 3441
      }
    ]
  },
  "minecraft:clay": {
    "states": [
      {
        "id": 3442,
        "default": true
      }
    ]
  },
  "minecraft:sugar_cane": {
    "properties": {
      "age": [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "10",
        "11",
        "12",
        "13",
        "14",
        "15"
      ]
    },
    "states": [
      {
        "properties": {
          "age": "0"
        },
        "id": 3443,
        "default": true
      },
      {
        "properties": {
          "age": "1"
        },
        "id": 3444
      },
      {
        "properties": {
          "age": "2"
        },
        "id": 3445
      },
      {
        "properties": {
          "age": "3"
        },
        "id": 3446
      },
      {
        "properties": {
          "age": "4"
        },
        "id": 3447
      },
      {
        "properties": {
          "age": "5"
        },
        "id": 3448
      },
      {
        "properties": {
          "age": "6"
        },
        "id": 3449
      },
      {
        "properties": {
          "age": "7"
        },
        "id": 3450
      },
      {
        "properties": {
          "age": "8"
        },
        "id": 3451
      },
      {
        "properties": {
          "age": "9"
        },
        "id": 3452
      },
      {
        "properties": {
          "age": "10"
        },
        "id": 3453
      },
      {
        "properties": {
          "age": "11"
        },
        "id": 3454
      },
      {
        "properties": {
          "age": "12"
        },
        "id": 3455
      },
      {
        "properties": {
          "age": "13"
        },
        "id": 3456
      },
      {
        "properties": {
          "age": "14"
        },
        "id": 3457
      },
      {
        "properties": {
          "age": "15"
        },
        "id": 3458
      }
    ]
  },
  "minecraft:jukebox": {
    "properties": {
      "has_record": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "has_record": "true"
        },
        "id": 3459
      },
      {
        "properties": {
          "has_record": "false"
        },
        "id": 3460,
        "default": true
      }
    ]
  },
  "minecraft:oak_fence": {
    "properties": {
      "east": [
        "true",
        "false"
      ],
      "north": [
        "true",
        "false"
      ],
      "south": [
        "true",
        "false"
      ],
      "waterlogged": [
        "true",
        "false"
      ],
      "west": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 3461
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 3462
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 3463
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 3464
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 3465
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 3466
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 3467
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 3468
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 3469
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 3470
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 3471
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 3472
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 3473
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 3474
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 3475
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 3476
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 3477
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 3478
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 3479
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 3480
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 3481
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 3482
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 3483
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 3484
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 3485
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 3486
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 3487
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 3488
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 3489
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 3490
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 3491
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 3492,
        "default": true
      }
    ]
  },
  "minecraft:pumpkin": {
    "states": [
      {
        "id": 3493,
        "default": true
      }
    ]
  },
  "minecraft:netherrack": {
    "states": [
      {
        "id": 3494,
        "default": true
      }
    ]
  },
  "minecraft:soul_sand": {
    "states": [
      {
        "id": 3495,
        "default": true
      }
    ]
  },
  "minecraft:glowstone": {
    "states": [
      {
        "id": 3496,
        "default": true
      }
    ]
  },
  "minecraft:nether_portal": {
    "properties": {
      "axis": [
        "x",
        "z"
      ]
    },
    "states": [
      {
        "properties": {
          "axis": "x"
        },
        "id": 3497,
        "default": true
      },
      {
        "properties": {
          "axis": "z"
        },
        "id": 3498
      }
    ]
  },
  "minecraft:carved_pumpkin": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north"
        },
        "id": 3499,
        "default": true
      },
      {
        "properties": {
          "facing": "south"
        },
        "id": 3500
      },
      {
        "properties": {
          "facing": "west"
        },
        "id": 3501
      },
      {
        "properties": {
          "facing": "east"
        },
        "id": 3502
      }
    ]
  },
  "minecraft:jack_o_lantern": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north"
        },
        "id": 3503,
        "default": true
      },
      {
        "properties": {
          "facing": "south"
        },
        "id": 3504
      },
      {
        "properties": {
          "facing": "west"
        },
        "id": 3505
      },
      {
        "properties": {
          "facing": "east"
        },
        "id": 3506
      }
    ]
  },
  "minecraft:cake": {
    "properties": {
      "bites": [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6"
      ]
    },
    "states": [
      {
        "properties": {
          "bites": "0"
        },
        "id": 3507,
        "default": true
      },
      {
        "properties": {
          "bites": "1"
        },
        "id": 3508
      },
      {
        "properties": {
          "bites": "2"
        },
        "id": 3509
      },
      {
        "properties": {
          "bites": "3"
        },
        "id": 3510
      },
      {
        "properties": {
          "bites": "4"
        },
        "id": 3511
      },
      {
        "properties": {
          "bites": "5"
        },
        "id": 3512
      },
      {
        "properties": {
          "bites": "6"
        },
        "id": 3513
      }
    ]
  },
  "minecraft:repeater": {
    "properties": {
      "delay": [
        "1",
        "2",
        "3",
        "4"
      ],
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "locked": [
        "true",
        "false"
      ],
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "delay": "1",
          "facing": "north",
          "locked": "true",
          "powered": "true"
        },
        "id": 3514
      },
      {
        "properties": {
          "delay": "1",
          "facing": "north",
          "locked": "true",
          "powered": "false"
        },
        "id": 3515
      },
      {
        "properties": {
          "delay": "1",
          "facing": "north",
          "locked": "false",
          "powered": "true"
        },
        "id": 3516
      },
      {
        "properties": {
          "delay": "1",
          "facing": "north",
          "locked": "false",
          "powered": "false"
        },
        "id": 3517,
        "default": true
      },
      {
        "properties": {
          "delay": "1",
          "facing": "south",
          "locked": "true",
          "powered": "true"
        },
        "id": 3518
      },
      {
        "properties": {
          "delay": "1",
          "facing": "south",
          "locked": "true",
          "powered": "false"
        },
        "id": 3519
      },
      {
        "properties": {
          "delay": "1",
          "facing": "south",
          "locked": "false",
          "powered": "true"
        },
        "id": 3520
      },
      {
        "properties": {
          "delay": "1",
          "facing": "south",
          "locked": "false",
          "powered": "false"
        },
        "id": 3521
      },
      {
        "properties": {
          "delay": "1",
          "facing": "west",
          "locked": "true",
          "powered": "true"
        },
        "id": 3522
      },
      {
        "properties": {
          "delay": "1",
          "facing": "west",
          "locked": "true",
          "powered": "false"
        },
        "id": 3523
      },
      {
        "properties": {
          "delay": "1",
          "facing": "west",
          "locked": "false",
          "powered": "true"
        },
        "id": 3524
      },
      {
        "properties": {
          "delay": "1",
          "facing": "west",
          "locked": "false",
          "powered": "false"
        },
        "id": 3525
      },
      {
        "properties": {
          "delay": "1",
          "facing": "east",
          "locked": "true",
          "powered": "true"
        },
        "id": 3526
      },
      {
        "properties": {
          "delay": "1",
          "facing": "east",
          "locked": "true",
          "powered": "false"
        },
        "id": 3527
      },
      {
        "properties": {
          "delay": "1",
          "facing": "east",
          "locked": "false",
          "powered": "true"
        },
        "id": 3528
      },
      {
        "properties": {
          "delay": "1",
          "facing": "east",
          "locked": "false",
          "powered": "false"
        },
        "id": 3529
      },
      {
        "properties": {
          "delay": "2",
          "facing": "north",
          "locked": "true",
          "powered": "true"
        },
        "id": 3530
      },
      {
        "properties": {
          "delay": "2",
          "facing": "north",
          "locked": "true",
          "powered": "false"
        },
        "id": 3531
      },
      {
        "properties": {
          "delay": "2",
          "facing": "north",
          "locked": "false",
          "powered": "true"
        },
        "id": 3532
      },
      {
        "properties": {
          "delay": "2",
          "facing": "north",
          "locked": "false",
          "powered": "false"
        },
        "id": 3533
      },
      {
        "properties": {
          "delay": "2",
          "facing": "south",
          "locked": "true",
          "powered": "true"
        },
        "id": 3534
      },
      {
        "properties": {
          "delay": "2",
          "facing": "south",
          "locked": "true",
          "powered": "false"
        },
        "id": 3535
      },
      {
        "properties": {
          "delay": "2",
          "facing": "south",
          "locked": "false",
          "powered": "true"
        },
        "id": 3536
      },
      {
        "properties": {
          "delay": "2",
          "facing": "south",
          "locked": "false",
          "powered": "false"
        },
        "id": 3537
      },
      {
        "properties": {
          "delay": "2",
          "facing": "west",
          "locked": "true",
          "powered": "true"
        },
        "id": 3538
      },
      {
        "properties": {
          "delay": "2",
          "facing": "west",
          "locked": "true",
          "powered": "false"
        },
        "id": 3539
      },
      {
        "properties": {
          "delay": "2",
          "facing": "west",
          "locked": "false",
          "powered": "true"
        },
        "id": 3540
      },
      {
        "properties": {
          "delay": "2",
          "facing": "west",
          "locked": "false",
          "powered": "false"
        },
        "id": 3541
      },
      {
        "properties": {
          "delay": "2",
          "facing": "east",
          "locked": "true",
          "powered": "true"
        },
        "id": 3542
      },
      {
        "properties": {
          "delay": "2",
          "facing": "east",
          "locked": "true",
          "powered": "false"
        },
        "id": 3543
      },
      {
        "properties": {
          "delay": "2",
          "facing": "east",
          "locked": "false",
          "powered": "true"
        },
        "id": 3544
      },
      {
        "properties": {
          "delay": "2",
          "facing": "east",
          "locked": "false",
          "powered": "false"
        },
        "id": 3545
      },
      {
        "properties": {
          "delay": "3",
          "facing": "north",
          "locked": "true",
          "powered": "true"
        },
        "id": 3546
      },
      {
        "properties": {
          "delay": "3",
          "facing": "north",
          "locked": "true",
          "powered": "false"
        },
        "id": 3547
      },
      {
        "properties": {
          "delay": "3",
          "facing": "north",
          "locked": "false",
          "powered": "true"
        },
        "id": 3548
      },
      {
        "properties": {
          "delay": "3",
          "facing": "north",
          "locked": "false",
          "powered": "false"
        },
        "id": 3549
      },
      {
        "properties": {
          "delay": "3",
          "facing": "south",
          "locked": "true",
          "powered": "true"
        },
        "id": 3550
      },
      {
        "properties": {
          "delay": "3",
          "facing": "south",
          "locked": "true",
          "powered": "false"
        },
        "id": 3551
      },
      {
        "properties": {
          "delay": "3",
          "facing": "south",
          "locked": "false",
          "powered": "true"
        },
        "id": 3552
      },
      {
        "properties": {
          "delay": "3",
          "facing": "south",
          "locked": "false",
          "powered": "false"
        },
        "id": 3553
      },
      {
        "properties": {
          "delay": "3",
          "facing": "west",
          "locked": "true",
          "powered": "true"
        },
        "id": 3554
      },
      {
        "properties": {
          "delay": "3",
          "facing": "west",
          "locked": "true",
          "powered": "false"
        },
        "id": 3555
      },
      {
        "properties": {
          "delay": "3",
          "facing": "west",
          "locked": "false",
          "powered": "true"
        },
        "id": 3556
      },
      {
        "properties": {
          "delay": "3",
          "facing": "west",
          "locked": "false",
          "powered": "false"
        },
        "id": 3557
      },
      {
        "properties": {
          "delay": "3",
          "facing": "east",
          "locked": "true",
          "powered": "true"
        },
        "id": 3558
      },
      {
        "properties": {
          "delay": "3",
          "facing": "east",
          "locked": "true",
          "powered": "false"
        },
        "id": 3559
      },
      {
        "properties": {
          "delay": "3",
          "facing": "east",
          "locked": "false",
          "powered": "true"
        },
        "id": 3560
      },
      {
        "properties": {
          "delay": "3",
          "facing": "east",
          "locked": "false",
          "powered": "false"
        },
        "id": 3561
      },
      {
        "properties": {
          "delay": "4",
          "facing": "north",
          "locked": "true",
          "powered": "true"
        },
        "id": 3562
      },
      {
        "properties": {
          "delay": "4",
          "facing": "north",
          "locked": "true",
          "powered": "false"
        },
        "id": 3563
      },
      {
        "properties": {
          "delay": "4",
          "facing": "north",
          "locked": "false",
          "powered": "true"
        },
        "id": 3564
      },
      {
        "properties": {
          "delay": "4",
          "facing": "north",
          "locked": "false",
          "powered": "false"
        },
        "id": 3565
      },
      {
        "properties": {
          "delay": "4",
          "facing": "south",
          "locked": "true",
          "powered": "true"
        },
        "id": 3566
      },
      {
        "properties": {
          "delay": "4",
          "facing": "south",
          "locked": "true",
          "powered": "false"
        },
        "id": 3567
      },
      {
        "properties": {
          "delay": "4",
          "facing": "south",
          "locked": "false",
          "powered": "true"
        },
        "id": 3568
      },
      {
        "properties": {
          "delay": "4",
          "facing": "south",
          "locked": "false",
          "powered": "false"
        },
        "id": 3569
      },
      {
        "properties": {
          "delay": "4",
          "facing": "west",
          "locked": "true",
          "powered": "true"
        },
        "id": 3570
      },
      {
        "properties": {
          "delay": "4",
          "facing": "west",
          "locked": "true",
          "powered": "false"
        },
        "id": 3571
      },
      {
        "properties": {
          "delay": "4",
          "facing": "west",
          "locked": "false",
          "powered": "true"
        },
        "id": 3572
      },
      {
        "properties": {
          "delay": "4",
          "facing": "west",
          "locked": "false",
          "powered": "false"
        },
        "id": 3573
      },
      {
        "properties": {
          "delay": "4",
          "facing": "east",
          "locked": "true",
          "powered": "true"
        },
        "id": 3574
      },
      {
        "properties": {
          "delay": "4",
          "facing": "east",
          "locked": "true",
          "powered": "false"
        },
        "id": 3575
      },
      {
        "properties": {
          "delay": "4",
          "facing": "east",
          "locked": "false",
          "powered": "true"
        },
        "id": 3576
      },
      {
        "properties": {
          "delay": "4",
          "facing": "east",
          "locked": "false",
          "powered": "false"
        },
        "id": 3577
      }
    ]
  },
  "minecraft:white_stained_glass": {
    "states": [
      {
        "id": 3578,
        "default": true
      }
    ]
  },
  "minecraft:orange_stained_glass": {
    "states": [
      {
        "id": 3579,
        "default": true
      }
    ]
  },
  "minecraft:magenta_stained_glass": {
    "states": [
      {
        "id": 3580,
        "default": true
      }
    ]
  },
  "minecraft:light_blue_stained_glass": {
    "states": [
      {
        "id": 3581,
        "default": true
      }
    ]
  },
  "minecraft:yellow_stained_glass": {
    "states": [
      {
        "id": 3582,
        "default": true
      }
    ]
  },
  "minecraft:lime_stained_glass": {
    "states": [
      {
        "id": 3583,
        "default": true
      }
    ]
  },
  "minecraft:pink_stained_glass": {
    "states": [
      {
        "id": 3584,
        "default": true
      }
    ]
  },
  "minecraft:gray_stained_glass": {
    "states": [
      {
        "id": 3585,
        "default": true
      }
    ]
  },
  "minecraft:light_gray_stained_glass": {
    "states": [
      {
        "id": 3586,
        "default": true
      }
    ]
  },
  "minecraft:cyan_stained_glass": {
    "states": [
      {
        "id": 3587,
        "default": true
      }
    ]
  },
  "minecraft:purple_stained_glass": {
    "states": [
      {
        "id": 3588,
        "default": true
      }
    ]
  },
  "minecraft:blue_stained_glass": {
    "states": [
      {
        "id": 3589,
        "default": true
      }
    ]
  },
  "minecraft:brown_stained_glass": {
    "states": [
      {
        "id": 3590,
        "default": true
      }
    ]
  },
  "minecraft:green_stained_glass": {
    "states": [
      {
        "id": 3591,
        "default": true
      }
    ]
  },
  "minecraft:red_stained_glass": {
    "states": [
      {
        "id": 3592,
        "default": true
      }
    ]
  },
  "minecraft:black_stained_glass": {
    "states": [
      {
        "id": 3593,
        "default": true
      }
    ]
  },
  "minecraft:oak_trapdoor": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "open": [
        "true",
        "false"
      ],
      "powered": [
        "true",
        "false"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3594
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3595
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3596
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3597
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3598
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3599
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3600
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3601
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3602
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3603
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3604
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3605
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3606
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3607
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3608
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3609,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3610
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3611
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3612
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3613
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3614
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3615
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3616
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3617
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3618
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3619
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3620
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3621
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3622
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3623
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3624
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3625
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3626
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3627
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3628
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3629
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3630
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3631
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3632
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3633
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3634
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3635
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3636
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3637
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3638
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3639
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3640
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3641
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3642
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3643
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3644
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3645
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3646
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3647
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3648
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3649
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3650
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3651
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3652
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3653
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3654
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3655
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3656
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3657
      }
    ]
  },
  "minecraft:spruce_trapdoor": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "open": [
        "true",
        "false"
      ],
      "powered": [
        "true",
        "false"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3658
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3659
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3660
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3661
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3662
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3663
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3664
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3665
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3666
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3667
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3668
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3669
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3670
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3671
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3672
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3673,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3674
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3675
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3676
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3677
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3678
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3679
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3680
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3681
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3682
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3683
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3684
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3685
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3686
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3687
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3688
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3689
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3690
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3691
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3692
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3693
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3694
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3695
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3696
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3697
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3698
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3699
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3700
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3701
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3702
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3703
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3704
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3705
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3706
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3707
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3708
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3709
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3710
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3711
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3712
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3713
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3714
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3715
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3716
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3717
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3718
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3719
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3720
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3721
      }
    ]
  },
  "minecraft:birch_trapdoor": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "open": [
        "true",
        "false"
      ],
      "powered": [
        "true",
        "false"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3722
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3723
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3724
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3725
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3726
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3727
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3728
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3729
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3730
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3731
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3732
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3733
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3734
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3735
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3736
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3737,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3738
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3739
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3740
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3741
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3742
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3743
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3744
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3745
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3746
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3747
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3748
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3749
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3750
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3751
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3752
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3753
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3754
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3755
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3756
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3757
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3758
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3759
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3760
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3761
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3762
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3763
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3764
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3765
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3766
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3767
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3768
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3769
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3770
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3771
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3772
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3773
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3774
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3775
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3776
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3777
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3778
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3779
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3780
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3781
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3782
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3783
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3784
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3785
      }
    ]
  },
  "minecraft:jungle_trapdoor": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "open": [
        "true",
        "false"
      ],
      "powered": [
        "true",
        "false"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3786
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3787
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3788
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3789
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3790
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3791
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3792
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3793
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3794
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3795
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3796
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3797
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3798
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3799
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3800
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3801,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3802
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3803
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3804
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3805
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3806
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3807
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3808
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3809
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3810
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3811
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3812
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3813
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3814
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3815
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3816
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3817
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3818
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3819
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3820
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3821
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3822
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3823
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3824
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3825
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3826
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3827
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3828
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3829
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3830
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3831
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3832
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3833
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3834
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3835
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3836
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3837
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3838
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3839
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3840
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3841
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3842
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3843
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3844
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3845
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3846
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3847
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3848
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3849
      }
    ]
  },
  "minecraft:acacia_trapdoor": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "open": [
        "true",
        "false"
      ],
      "powered": [
        "true",
        "false"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3850
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3851
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3852
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3853
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3854
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3855
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3856
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3857
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3858
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3859
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3860
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3861
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3862
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3863
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3864
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3865,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3866
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3867
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3868
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3869
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3870
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3871
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3872
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3873
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3874
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3875
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3876
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3877
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3878
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3879
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3880
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3881
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3882
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3883
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3884
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3885
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3886
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3887
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3888
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3889
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3890
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3891
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3892
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3893
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3894
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3895
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3896
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3897
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3898
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3899
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3900
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3901
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3902
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3903
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3904
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3905
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3906
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3907
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3908
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3909
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3910
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3911
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3912
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3913
      }
    ]
  },
  "minecraft:dark_oak_trapdoor": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "open": [
        "true",
        "false"
      ],
      "powered": [
        "true",
        "false"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3914
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3915
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3916
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3917
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3918
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3919
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3920
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3921
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3922
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3923
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3924
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3925
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3926
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3927
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3928
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3929,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3930
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3931
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3932
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3933
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3934
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3935
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3936
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3937
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3938
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3939
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3940
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3941
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3942
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3943
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3944
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3945
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3946
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3947
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3948
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3949
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3950
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3951
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3952
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3953
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3954
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3955
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3956
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3957
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3958
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3959
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3960
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3961
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3962
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3963
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3964
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3965
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3966
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3967
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3968
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3969
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3970
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3971
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3972
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "true",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3973
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "true"
        },
        "id": 3974
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "true",
          "waterlogged": "false"
        },
        "id": 3975
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "true"
        },
        "id": 3976
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "open": "false",
          "powered": "false",
          "waterlogged": "false"
        },
        "id": 3977
      }
    ]
  },
  "minecraft:infested_stone": {
    "states": [
      {
        "id": 3978,
        "default": true
      }
    ]
  },
  "minecraft:infested_cobblestone": {
    "states": [
      {
        "id": 3979,
        "default": true
      }
    ]
  },
  "minecraft:infested_stone_bricks": {
    "states": [
      {
        "id": 3980,
        "default": true
      }
    ]
  },
  "minecraft:infested_mossy_stone_bricks": {
    "states": [
      {
        "id": 3981,
        "default": true
      }
    ]
  },
  "minecraft:infested_cracked_stone_bricks": {
    "states": [
      {
        "id": 3982,
        "default": true
      }
    ]
  },
  "minecraft:infested_chiseled_stone_bricks": {
    "states": [
      {
        "id": 3983,
        "default": true
      }
    ]
  },
  "minecraft:stone_bricks": {
    "states": [
      {
        "id": 3984,
        "default": true
      }
    ]
  },
  "minecraft:mossy_stone_bricks": {
    "states": [
      {
        "id": 3985,
        "default": true
      }
    ]
  },
  "minecraft:cracked_stone_bricks": {
    "states": [
      {
        "id": 3986,
        "default": true
      }
    ]
  },
  "minecraft:chiseled_stone_bricks": {
    "states": [
      {
        "id": 3987,
        "default": true
      }
    ]
  },
  "minecraft:brown_mushroom_block": {
    "properties": {
      "down": [
        "true",
        "false"
      ],
      "east": [
        "true",
        "false"
      ],
      "north": [
        "true",
        "false"
      ],
      "south": [
        "true",
        "false"
      ],
      "up": [
        "true",
        "false"
      ],
      "west": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 3988,
        "default": true
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 3989
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 3990
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 3991
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 3992
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 3993
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 3994
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 3995
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 3996
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 3997
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 3998
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 3999
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4000
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4001
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4002
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4003
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4004
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4005
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4006
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4007
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4008
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4009
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4010
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4011
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4012
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4013
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4014
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4015
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4016
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4017
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4018
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4019
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4020
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4021
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4022
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4023
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4024
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4025
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4026
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4027
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4028
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4029
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4030
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4031
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4032
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4033
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4034
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4035
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4036
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4037
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4038
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4039
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4040
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4041
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4042
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4043
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4044
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4045
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4046
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4047
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4048
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4049
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4050
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4051
      }
    ]
  },
  "minecraft:red_mushroom_block": {
    "properties": {
      "down": [
        "true",
        "false"
      ],
      "east": [
        "true",
        "false"
      ],
      "north": [
        "true",
        "false"
      ],
      "south": [
        "true",
        "false"
      ],
      "up": [
        "true",
        "false"
      ],
      "west": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4052,
        "default": true
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4053
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4054
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4055
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4056
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4057
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4058
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4059
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4060
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4061
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4062
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4063
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4064
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4065
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4066
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4067
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4068
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4069
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4070
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4071
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4072
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4073
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4074
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4075
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4076
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4077
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4078
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4079
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4080
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4081
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4082
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4083
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4084
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4085
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4086
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4087
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4088
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4089
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4090
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4091
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4092
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4093
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4094
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4095
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4096
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4097
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4098
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4099
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4100
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4101
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4102
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4103
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4104
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4105
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4106
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4107
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4108
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4109
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4110
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4111
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4112
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4113
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4114
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4115
      }
    ]
  },
  "minecraft:mushroom_stem": {
    "properties": {
      "down": [
        "true",
        "false"
      ],
      "east": [
        "true",
        "false"
      ],
      "north": [
        "true",
        "false"
      ],
      "south": [
        "true",
        "false"
      ],
      "up": [
        "true",
        "false"
      ],
      "west": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4116,
        "default": true
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4117
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4118
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4119
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4120
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4121
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4122
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4123
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4124
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4125
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4126
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4127
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4128
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4129
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4130
      },
      {
        "properties": {
          "down": "true",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4131
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4132
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4133
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4134
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4135
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4136
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4137
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4138
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4139
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4140
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4141
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4142
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4143
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4144
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4145
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4146
      },
      {
        "properties": {
          "down": "true",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4147
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4148
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4149
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4150
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4151
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4152
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4153
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4154
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4155
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4156
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4157
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4158
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4159
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4160
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4161
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4162
      },
      {
        "properties": {
          "down": "false",
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4163
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4164
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4165
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4166
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4167
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4168
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4169
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4170
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4171
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4172
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4173
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4174
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4175
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4176
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4177
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4178
      },
      {
        "properties": {
          "down": "false",
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4179
      }
    ]
  },
  "minecraft:iron_bars": {
    "properties": {
      "east": [
        "true",
        "false"
      ],
      "north": [
        "true",
        "false"
      ],
      "south": [
        "true",
        "false"
      ],
      "waterlogged": [
        "true",
        "false"
      ],
      "west": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4180
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4181
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4182
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4183
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4184
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4185
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4186
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4187
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4188
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4189
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4190
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4191
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4192
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4193
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4194
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4195
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4196
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4197
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4198
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4199
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4200
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4201
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4202
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4203
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4204
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4205
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4206
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4207
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4208
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4209
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4210
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4211,
        "default": true
      }
    ]
  },
  "minecraft:glass_pane": {
    "properties": {
      "east": [
        "true",
        "false"
      ],
      "north": [
        "true",
        "false"
      ],
      "south": [
        "true",
        "false"
      ],
      "waterlogged": [
        "true",
        "false"
      ],
      "west": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4212
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4213
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4214
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4215
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4216
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4217
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4218
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4219
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4220
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4221
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4222
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4223
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4224
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4225
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4226
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4227
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4228
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4229
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4230
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4231
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4232
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4233
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4234
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4235
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4236
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4237
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4238
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4239
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4240
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4241
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4242
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4243,
        "default": true
      }
    ]
  },
  "minecraft:melon": {
    "states": [
      {
        "id": 4244,
        "default": true
      }
    ]
  },
  "minecraft:attached_pumpkin_stem": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north"
        },
        "id": 4245,
        "default": true
      },
      {
        "properties": {
          "facing": "south"
        },
        "id": 4246
      },
      {
        "properties": {
          "facing": "west"
        },
        "id": 4247
      },
      {
        "properties": {
          "facing": "east"
        },
        "id": 4248
      }
    ]
  },
  "minecraft:attached_melon_stem": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north"
        },
        "id": 4249,
        "default": true
      },
      {
        "properties": {
          "facing": "south"
        },
        "id": 4250
      },
      {
        "properties": {
          "facing": "west"
        },
        "id": 4251
      },
      {
        "properties": {
          "facing": "east"
        },
        "id": 4252
      }
    ]
  },
  "minecraft:pumpkin_stem": {
    "properties": {
      "age": [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7"
      ]
    },
    "states": [
      {
        "properties": {
          "age": "0"
        },
        "id": 4253,
        "default": true
      },
      {
        "properties": {
          "age": "1"
        },
        "id": 4254
      },
      {
        "properties": {
          "age": "2"
        },
        "id": 4255
      },
      {
        "properties": {
          "age": "3"
        },
        "id": 4256
      },
      {
        "properties": {
          "age": "4"
        },
        "id": 4257
      },
      {
        "properties": {
          "age": "5"
        },
        "id": 4258
      },
      {
        "properties": {
          "age": "6"
        },
        "id": 4259
      },
      {
        "properties": {
          "age": "7"
        },
        "id": 4260
      }
    ]
  },
  "minecraft:melon_stem": {
    "properties": {
      "age": [
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7"
      ]
    },
    "states": [
      {
        "properties": {
          "age": "0"
        },
        "id": 4261,
        "default": true
      },
      {
        "properties": {
          "age": "1"
        },
        "id": 4262
      },
      {
        "properties": {
          "age": "2"
        },
        "id": 4263
      },
      {
        "properties": {
          "age": "3"
        },
        "id": 4264
      },
      {
        "properties": {
          "age": "4"
        },
        "id": 4265
      },
      {
        "properties": {
          "age": "5"
        },
        "id": 4266
      },
      {
        "properties": {
          "age": "6"
        },
        "id": 4267
      },
      {
        "properties": {
          "age": "7"
        },
        "id": 4268
      }
    ]
  },
  "minecraft:vine": {
    "properties": {
      "east": [
        "true",
        "false"
      ],
      "north": [
        "true",
        "false"
      ],
      "south": [
        "true",
        "false"
      ],
      "up": [
        "true",
        "false"
      ],
      "west": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4269
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4270
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4271
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4272
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4273
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4274
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4275
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4276
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4277
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4278
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4279
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4280
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4281
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4282
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4283
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4284
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4285
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4286
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4287
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4288
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4289
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4290
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4291
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4292
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "true"
        },
        "id": 4293
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "true",
          "west": "false"
        },
        "id": 4294
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "true"
        },
        "id": 4295
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "up": "false",
          "west": "false"
        },
        "id": 4296
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "true"
        },
        "id": 4297
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "true",
          "west": "false"
        },
        "id": 4298
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "true"
        },
        "id": 4299
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "up": "false",
          "west": "false"
        },
        "id": 4300,
        "default": true
      }
    ]
  },
  "minecraft:oak_fence_gate": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "in_wall": [
        "true",
        "false"
      ],
      "open": [
        "true",
        "false"
      ],
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "in_wall": "true",
          "open": "true",
          "powered": "true"
        },
        "id": 4301
      },
      {
        "properties": {
          "facing": "north",
          "in_wall": "true",
          "open": "true",
          "powered": "false"
        },
        "id": 4302
      },
      {
        "properties": {
          "facing": "north",
          "in_wall": "true",
          "open": "false",
          "powered": "true"
        },
        "id": 4303
      },
      {
        "properties": {
          "facing": "north",
          "in_wall": "true",
          "open": "false",
          "powered": "false"
        },
        "id": 4304
      },
      {
        "properties": {
          "facing": "north",
          "in_wall": "false",
          "open": "true",
          "powered": "true"
        },
        "id": 4305
      },
      {
        "properties": {
          "facing": "north",
          "in_wall": "false",
          "open": "true",
          "powered": "false"
        },
        "id": 4306
      },
      {
        "properties": {
          "facing": "north",
          "in_wall": "false",
          "open": "false",
          "powered": "true"
        },
        "id": 4307
      },
      {
        "properties": {
          "facing": "north",
          "in_wall": "false",
          "open": "false",
          "powered": "false"
        },
        "id": 4308,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "in_wall": "true",
          "open": "true",
          "powered": "true"
        },
        "id": 4309
      },
      {
        "properties": {
          "facing": "south",
          "in_wall": "true",
          "open": "true",
          "powered": "false"
        },
        "id": 4310
      },
      {
        "properties": {
          "facing": "south",
          "in_wall": "true",
          "open": "false",
          "powered": "true"
        },
        "id": 4311
      },
      {
        "properties": {
          "facing": "south",
          "in_wall": "true",
          "open": "false",
          "powered": "false"
        },
        "id": 4312
      },
      {
        "properties": {
          "facing": "south",
          "in_wall": "false",
          "open": "true",
          "powered": "true"
        },
        "id": 4313
      },
      {
        "properties": {
          "facing": "south",
          "in_wall": "false",
          "open": "true",
          "powered": "false"
        },
        "id": 4314
      },
      {
        "properties": {
          "facing": "south",
          "in_wall": "false",
          "open": "false",
          "powered": "true"
        },
        "id": 4315
      },
      {
        "properties": {
          "facing": "south",
          "in_wall": "false",
          "open": "false",
          "powered": "false"
        },
        "id": 4316
      },
      {
        "properties": {
          "facing": "west",
          "in_wall": "true",
          "open": "true",
          "powered": "true"
        },
        "id": 4317
      },
      {
        "properties": {
          "facing": "west",
          "in_wall": "true",
          "open": "true",
          "powered": "false"
        },
        "id": 4318
      },
      {
        "properties": {
          "facing": "west",
          "in_wall": "true",
          "open": "false",
          "powered": "true"
        },
        "id": 4319
      },
      {
        "properties": {
          "facing": "west",
          "in_wall": "true",
          "open": "false",
          "powered": "false"
        },
        "id": 4320
      },
      {
        "properties": {
          "facing": "west",
          "in_wall": "false",
          "open": "true",
          "powered": "true"
        },
        "id": 4321
      },
      {
        "properties": {
          "facing": "west",
          "in_wall": "false",
          "open": "true",
          "powered": "false"
        },
        "id": 4322
      },
      {
        "properties": {
          "facing": "west",
          "in_wall": "false",
          "open": "false",
          "powered": "true"
        },
        "id": 4323
      },
      {
        "properties": {
          "facing": "west",
          "in_wall": "false",
          "open": "false",
          "powered": "false"
        },
        "id": 4324
      },
      {
        "properties": {
          "facing": "east",
          "in_wall": "true",
          "open": "true",
          "powered": "true"
        },
        "id": 4325
      },
      {
        "properties": {
          "facing": "east",
          "in_wall": "true",
          "open": "true",
          "powered": "false"
        },
        "id": 4326
      },
      {
        "properties": {
          "facing": "east",
          "in_wall": "true",
          "open": "false",
          "powered": "true"
        },
        "id": 4327
      },
      {
        "properties": {
          "facing": "east",
          "in_wall": "true",
          "open": "false",
          "powered": "false"
        },
        "id": 4328
      },
      {
        "properties": {
          "facing": "east",
          "in_wall": "false",
          "open": "true",
          "powered": "true"
        },
        "id": 4329
      },
      {
        "properties": {
          "facing": "east",
          "in_wall": "false",
          "open": "true",
          "powered": "false"
        },
        "id": 4330
      },
      {
        "properties": {
          "facing": "east",
          "in_wall": "false",
          "open": "false",
          "powered": "true"
        },
        "id": 4331
      },
      {
        "properties": {
          "facing": "east",
          "in_wall": "false",
          "open": "false",
          "powered": "false"
        },
        "id": 4332
      }
    ]
  },
  "minecraft:brick_stairs": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "shape": [
        "straight",
        "inner_left",
        "inner_right",
        "outer_left",
        "outer_right"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4333
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4334
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4335
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4336
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4337
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4338
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4339
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4340
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4341
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4342
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4343
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4344,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4345
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4346
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4347
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4348
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4349
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4350
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4351
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4352
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4353
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4354
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4355
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4356
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4357
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4358
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4359
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4360
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4361
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4362
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4363
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4364
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4365
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4366
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4367
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4368
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4369
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4370
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4371
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4372
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4373
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4374
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4375
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4376
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4377
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4378
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4379
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4380
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4381
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4382
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4383
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4384
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4385
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4386
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4387
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4388
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4389
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4390
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4391
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4392
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4393
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4394
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4395
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4396
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4397
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4398
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4399
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4400
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4401
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4402
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4403
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4404
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4405
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4406
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4407
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4408
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4409
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4410
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4411
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4412
      }
    ]
  },
  "minecraft:stone_brick_stairs": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "shape": [
        "straight",
        "inner_left",
        "inner_right",
        "outer_left",
        "outer_right"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4413
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4414
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4415
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4416
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4417
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4418
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4419
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4420
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4421
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4422
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4423
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4424,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4425
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4426
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4427
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4428
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4429
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4430
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4431
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4432
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4433
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4434
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4435
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4436
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4437
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4438
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4439
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4440
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4441
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4442
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4443
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4444
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4445
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4446
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4447
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4448
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4449
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4450
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4451
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4452
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4453
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4454
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4455
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4456
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4457
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4458
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4459
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4460
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4461
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4462
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4463
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4464
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4465
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4466
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4467
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4468
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4469
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4470
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4471
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4472
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4473
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4474
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4475
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4476
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4477
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4478
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4479
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4480
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4481
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4482
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4483
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4484
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4485
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4486
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4487
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4488
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4489
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4490
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4491
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4492
      }
    ]
  },
  "minecraft:mycelium": {
    "properties": {
      "snowy": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "snowy": "true"
        },
        "id": 4493
      },
      {
        "properties": {
          "snowy": "false"
        },
        "id": 4494,
        "default": true
      }
    ]
  },
  "minecraft:lily_pad": {
    "states": [
      {
        "id": 4495,
        "default": true
      }
    ]
  },
  "minecraft:nether_bricks": {
    "states": [
      {
        "id": 4496,
        "default": true
      }
    ]
  },
  "minecraft:nether_brick_fence": {
    "properties": {
      "east": [
        "true",
        "false"
      ],
      "north": [
        "true",
        "false"
      ],
      "south": [
        "true",
        "false"
      ],
      "waterlogged": [
        "true",
        "false"
      ],
      "west": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4497
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4498
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4499
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4500
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4501
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4502
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4503
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4504
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4505
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4506
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4507
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4508
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4509
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4510
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4511
      },
      {
        "properties": {
          "east": "true",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4512
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4513
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4514
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4515
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4516
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4517
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4518
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4519
      },
      {
        "properties": {
          "east": "false",
          "north": "true",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4520
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4521
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4522
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4523
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4524
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 4525
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 4526
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 4527
      },
      {
        "properties": {
          "east": "false",
          "north": "false",
          "south": "false",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 4528,
        "default": true
      }
    ]
  },
  "minecraft:nether_brick_stairs": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "shape": [
        "straight",
        "inner_left",
        "inner_right",
        "outer_left",
        "outer_right"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4529
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4530
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4531
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4532
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4533
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4534
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4535
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4536
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4537
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4538
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4539
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4540,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4541
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4542
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4543
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4544
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4545
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4546
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4547
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4548
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4549
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4550
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4551
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4552
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4553
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4554
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4555
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4556
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4557
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4558
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4559
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4560
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4561
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4562
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4563
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4564
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4565
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4566
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4567
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4568
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4569
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4570
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4571
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4572
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4573
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4574
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4575
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4576
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4577
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4578
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4579
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4580
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4581
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4582
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4583
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4584
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4585
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4586
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4587
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4588
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4589
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4590
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4591
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4592
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4593
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4594
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4595
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4596
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4597
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4598
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4599
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4600
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4601
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4602
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4603
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4604
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4605
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4606
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4607
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4608
      }
    ]
  },
  "minecraft:nether_wart": {
    "properties": {
      "age": [
        "0",
        "1",
        "2",
        "3"
      ]
    },
    "states": [
      {
        "properties": {
          "age": "0"
        },
        "id": 4609,
        "default": true
      },
      {
        "properties": {
          "age": "1"
        },
        "id": 4610
      },
      {
        "properties": {
          "age": "2"
        },
        "id": 4611
      },
      {
        "properties": {
          "age": "3"
        },
        "id": 4612
      }
    ]
  },
  "minecraft:enchanting_table": {
    "states": [
      {
        "id": 4613,
        "default": true
      }
    ]
  },
  "minecraft:brewing_stand": {
    "properties": {
      "has_bottle_0": [
        "true",
        "false"
      ],
      "has_bottle_1": [
        "true",
        "false"
      ],
      "has_bottle_2": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "has_bottle_0": "true",
          "has_bottle_1": "true",
          "has_bottle_2": "true"
        },
        "id": 4614
      },
      {
        "properties": {
          "has_bottle_0": "true",
          "has_bottle_1": "true",
          "has_bottle_2": "false"
        },
        "id": 4615
      },
      {
        "properties": {
          "has_bottle_0": "true",
          "has_bottle_1": "false",
          "has_bottle_2": "true"
        },
        "id": 4616
      },
      {
        "properties": {
          "has_bottle_0": "true",
          "has_bottle_1": "false",
          "has_bottle_2": "false"
        },
        "id": 4617
      },
      {
        "properties": {
          "has_bottle_0": "false",
          "has_bottle_1": "true",
          "has_bottle_2": "true"
        },
        "id": 4618
      },
      {
        "properties": {
          "has_bottle_0": "false",
          "has_bottle_1": "true",
          "has_bottle_2": "false"
        },
        "id": 4619
      },
      {
        "properties": {
          "has_bottle_0": "false",
          "has_bottle_1": "false",
          "has_bottle_2": "true"
        },
        "id": 4620
      },
      {
        "properties": {
          "has_bottle_0": "false",
          "has_bottle_1": "false",
          "has_bottle_2": "false"
        },
        "id": 4621,
        "default": true
      }
    ]
  },
  "minecraft:cauldron": {
    "properties": {
      "level": [
        "0",
        "1",
        "2",
        "3"
      ]
    },
    "states": [
      {
        "properties": {
          "level": "0"
        },
        "id": 4622,
        "default": true
      },
      {
        "properties": {
          "level": "1"
        },
        "id": 4623
      },
      {
        "properties": {
          "level": "2"
        },
        "id": 4624
      },
      {
        "properties": {
          "level": "3"
        },
        "id": 4625
      }
    ]
  },
  "minecraft:end_portal": {
    "states": [
      {
        "id": 4626,
        "default": true
      }
    ]
  },
  "minecraft:end_portal_frame": {
    "properties": {
      "eye": [
        "true",
        "false"
      ],
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ]
    },
    "states": [
      {
        "properties": {
          "eye": "true",
          "facing": "north"
        },
        "id": 4627
      },
      {
        "properties": {
          "eye": "true",
          "facing": "south"
        },
        "id": 4628
      },
      {
        "properties": {
          "eye": "true",
          "facing": "west"
        },
        "id": 4629
      },
      {
        "properties": {
          "eye": "true",
          "facing": "east"
        },
        "id": 4630
      },
      {
        "properties": {
          "eye": "false",
          "facing": "north"
        },
        "id": 4631,
        "default": true
      },
      {
        "properties": {
          "eye": "false",
          "facing": "south"
        },
        "id": 4632
      },
      {
        "properties": {
          "eye": "false",
          "facing": "west"
        },
        "id": 4633
      },
      {
        "properties": {
          "eye": "false",
          "facing": "east"
        },
        "id": 4634
      }
    ]
  },
  "minecraft:end_stone": {
    "states": [
      {
        "id": 4635,
        "default": true
      }
    ]
  },
  "minecraft:dragon_egg": {
    "states": [
      {
        "id": 4636,
        "default": true
      }
    ]
  },
  "minecraft:redstone_lamp": {
    "properties": {
      "lit": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "lit": "true"
        },
        "id": 4637
      },
      {
        "properties": {
          "lit": "false"
        },
        "id": 4638,
        "default": true
      }
    ]
  },
  "minecraft:cocoa": {
    "properties": {
      "age": [
        "0",
        "1",
        "2"
      ],
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ]
    },
    "states": [
      {
        "properties": {
          "age": "0",
          "facing": "north"
        },
        "id": 4639,
        "default": true
      },
      {
        "properties": {
          "age": "0",
          "facing": "south"
        },
        "id": 4640
      },
      {
        "properties": {
          "age": "0",
          "facing": "west"
        },
        "id": 4641
      },
      {
        "properties": {
          "age": "0",
          "facing": "east"
        },
        "id": 4642
      },
      {
        "properties": {
          "age": "1",
          "facing": "north"
        },
        "id": 4643
      },
      {
        "properties": {
          "age": "1",
          "facing": "south"
        },
        "id": 4644
      },
      {
        "properties": {
          "age": "1",
          "facing": "west"
        },
        "id": 4645
      },
      {
        "properties": {
          "age": "1",
          "facing": "east"
        },
        "id": 4646
      },
      {
        "properties": {
          "age": "2",
          "facing": "north"
        },
        "id": 4647
      },
      {
        "properties": {
          "age": "2",
          "facing": "south"
        },
        "id": 4648
      },
      {
        "properties": {
          "age": "2",
          "facing": "west"
        },
        "id": 4649
      },
      {
        "properties": {
          "age": "2",
          "facing": "east"
        },
        "id": 4650
      }
    ]
  },
  "minecraft:sandstone_stairs": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "shape": [
        "straight",
        "inner_left",
        "inner_right",
        "outer_left",
        "outer_right"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4651
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4652
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4653
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4654
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4655
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4656
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4657
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4658
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4659
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4660
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4661
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4662,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4663
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4664
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4665
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4666
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4667
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4668
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4669
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4670
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4671
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4672
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4673
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4674
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4675
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4676
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4677
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4678
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4679
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4680
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4681
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4682
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4683
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4684
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4685
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4686
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4687
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4688
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4689
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4690
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4691
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4692
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4693
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4694
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4695
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4696
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4697
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4698
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4699
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4700
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4701
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4702
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4703
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4704
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4705
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4706
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4707
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4708
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4709
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4710
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4711
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4712
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4713
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4714
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4715
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4716
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4717
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4718
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4719
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4720
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4721
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4722
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4723
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4724
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4725
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4726
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4727
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4728
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4729
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4730
      }
    ]
  },
  "minecraft:emerald_ore": {
    "states": [
      {
        "id": 4731,
        "default": true
      }
    ]
  },
  "minecraft:ender_chest": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "waterlogged": "true"
        },
        "id": 4732
      },
      {
        "properties": {
          "facing": "north",
          "waterlogged": "false"
        },
        "id": 4733,
        "default": true
      },
      {
        "properties": {
          "facing": "south",
          "waterlogged": "true"
        },
        "id": 4734
      },
      {
        "properties": {
          "facing": "south",
          "waterlogged": "false"
        },
        "id": 4735
      },
      {
        "properties": {
          "facing": "west",
          "waterlogged": "true"
        },
        "id": 4736
      },
      {
        "properties": {
          "facing": "west",
          "waterlogged": "false"
        },
        "id": 4737
      },
      {
        "properties": {
          "facing": "east",
          "waterlogged": "true"
        },
        "id": 4738
      },
      {
        "properties": {
          "facing": "east",
          "waterlogged": "false"
        },
        "id": 4739
      }
    ]
  },
  "minecraft:tripwire_hook": {
    "properties": {
      "attached": [
        "true",
        "false"
      ],
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "powered": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "attached": "true",
          "facing": "north",
          "powered": "true"
        },
        "id": 4740
      },
      {
        "properties": {
          "attached": "true",
          "facing": "north",
          "powered": "false"
        },
        "id": 4741
      },
      {
        "properties": {
          "attached": "true",
          "facing": "south",
          "powered": "true"
        },
        "id": 4742
      },
      {
        "properties": {
          "attached": "true",
          "facing": "south",
          "powered": "false"
        },
        "id": 4743
      },
      {
        "properties": {
          "attached": "true",
          "facing": "west",
          "powered": "true"
        },
        "id": 4744
      },
      {
        "properties": {
          "attached": "true",
          "facing": "west",
          "powered": "false"
        },
        "id": 4745
      },
      {
        "properties": {
          "attached": "true",
          "facing": "east",
          "powered": "true"
        },
        "id": 4746
      },
      {
        "properties": {
          "attached": "true",
          "facing": "east",
          "powered": "false"
        },
        "id": 4747
      },
      {
        "properties": {
          "attached": "false",
          "facing": "north",
          "powered": "true"
        },
        "id": 4748
      },
      {
        "properties": {
          "attached": "false",
          "facing": "north",
          "powered": "false"
        },
        "id": 4749,
        "default": true
      },
      {
        "properties": {
          "attached": "false",
          "facing": "south",
          "powered": "true"
        },
        "id": 4750
      },
      {
        "properties": {
          "attached": "false",
          "facing": "south",
          "powered": "false"
        },
        "id": 4751
      },
      {
        "properties": {
          "attached": "false",
          "facing": "west",
          "powered": "true"
        },
        "id": 4752
      },
      {
        "properties": {
          "attached": "false",
          "facing": "west",
          "powered": "false"
        },
        "id": 4753
      },
      {
        "properties": {
          "attached": "false",
          "facing": "east",
          "powered": "true"
        },
        "id": 4754
      },
      {
        "properties": {
          "attached": "false",
          "facing": "east",
          "powered": "false"
        },
        "id": 4755
      }
    ]
  },
  "minecraft:tripwire": {
    "properties": {
      "attached": [
        "true",
        "false"
      ],
      "disarmed": [
        "true",
        "false"
      ],
      "east": [
        "true",
        "false"
      ],
      "north": [
        "true",
        "false"
      ],
      "powered": [
        "true",
        "false"
      ],
      "south": [
        "true",
        "false"
      ],
      "west": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4756
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4757
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4758
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4759
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4760
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4761
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4762
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4763
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4764
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4765
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4766
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4767
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4768
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4769
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4770
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4771
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4772
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4773
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4774
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4775
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4776
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4777
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4778
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4779
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4780
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4781
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4782
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4783
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4784
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4785
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4786
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4787
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4788
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4789
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4790
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4791
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4792
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4793
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4794
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4795
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4796
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4797
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4798
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4799
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4800
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4801
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4802
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4803
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4804
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4805
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4806
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4807
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4808
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4809
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4810
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4811
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4812
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4813
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4814
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4815
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4816
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4817
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4818
      },
      {
        "properties": {
          "attached": "true",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4819
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4820
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4821
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4822
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4823
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4824
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4825
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4826
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4827
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4828
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4829
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4830
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4831
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4832
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4833
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4834
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4835
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4836
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4837
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4838
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4839
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4840
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4841
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4842
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4843
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4844
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4845
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4846
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4847
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4848
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4849
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4850
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "true",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4851
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4852
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4853
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4854
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4855
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4856
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4857
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4858
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4859
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4860
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4861
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4862
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4863
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4864
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4865
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4866
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "true",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4867
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4868
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4869
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4870
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4871
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4872
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4873
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4874
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "true",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4875
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "true"
        },
        "id": 4876
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "true",
          "west": "false"
        },
        "id": 4877
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "true"
        },
        "id": 4878
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "true",
          "south": "false",
          "west": "false"
        },
        "id": 4879
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "true"
        },
        "id": 4880
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "true",
          "west": "false"
        },
        "id": 4881
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "true"
        },
        "id": 4882
      },
      {
        "properties": {
          "attached": "false",
          "disarmed": "false",
          "east": "false",
          "north": "false",
          "powered": "false",
          "south": "false",
          "west": "false"
        },
        "id": 4883,
        "default": true
      }
    ]
  },
  "minecraft:emerald_block": {
    "states": [
      {
        "id": 4884,
        "default": true
      }
    ]
  },
  "minecraft:spruce_stairs": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "shape": [
        "straight",
        "inner_left",
        "inner_right",
        "outer_left",
        "outer_right"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4885
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4886
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4887
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4888
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4889
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4890
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4891
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4892
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4893
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4894
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4895
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4896,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4897
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4898
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4899
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4900
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4901
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4902
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4903
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4904
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4905
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4906
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4907
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4908
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4909
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4910
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4911
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4912
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4913
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4914
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4915
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4916
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4917
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4918
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4919
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4920
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4921
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4922
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4923
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4924
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4925
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4926
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4927
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4928
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4929
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4930
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4931
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4932
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4933
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4934
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4935
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4936
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4937
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4938
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4939
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4940
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4941
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4942
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4943
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4944
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4945
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4946
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4947
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4948
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4949
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4950
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4951
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4952
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4953
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4954
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4955
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4956
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4957
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4958
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4959
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4960
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4961
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4962
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4963
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4964
      }
    ]
  },
  "minecraft:birch_stairs": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "shape": [
        "straight",
        "inner_left",
        "inner_right",
        "outer_left",
        "outer_right"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4965
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4966
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4967
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4968
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4969
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4970
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4971
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4972
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4973
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4974
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4975
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4976,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4977
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4978
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4979
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4980
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4981
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4982
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4983
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4984
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4985
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4986
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4987
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4988
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4989
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 4990
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 4991
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 4992
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 4993
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 4994
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 4995
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 4996
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 4997
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 4998
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 4999
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 5000
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 5001
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 5002
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 5003
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 5004
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 5005
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 5006
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 5007
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 5008
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 5009
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 5010
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 5011
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 5012
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 5013
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 5014
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 5015
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 5016
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 5017
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 5018
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 5019
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 5020
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 5021
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 5022
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 5023
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 5024
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 5025
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 5026
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 5027
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 5028
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 5029
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 5030
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 5031
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 5032
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 5033
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 5034
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 5035
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 5036
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 5037
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 5038
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 5039
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 5040
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 5041
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 5042
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 5043
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 5044
      }
    ]
  },
  "minecraft:jungle_stairs": {
    "properties": {
      "facing": [
        "north",
        "south",
        "west",
        "east"
      ],
      "half": [
        "top",
        "bottom"
      ],
      "shape": [
        "straight",
        "inner_left",
        "inner_right",
        "outer_left",
        "outer_right"
      ],
      "waterlogged": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 5045
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 5046
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 5047
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 5048
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 5049
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 5050
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 5051
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 5052
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 5053
      },
      {
        "properties": {
          "facing": "north",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 5054
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 5055
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 5056,
        "default": true
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 5057
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 5058
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 5059
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 5060
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 5061
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 5062
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 5063
      },
      {
        "properties": {
          "facing": "north",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 5064
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 5065
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 5066
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 5067
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 5068
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 5069
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 5070
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 5071
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 5072
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 5073
      },
      {
        "properties": {
          "facing": "south",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 5074
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 5075
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 5076
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 5077
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 5078
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 5079
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 5080
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 5081
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 5082
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 5083
      },
      {
        "properties": {
          "facing": "south",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 5084
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 5085
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 5086
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 5087
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 5088
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 5089
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 5090
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 5091
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 5092
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 5093
      },
      {
        "properties": {
          "facing": "west",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 5094
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 5095
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 5096
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 5097
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 5098
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 5099
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 5100
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 5101
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 5102
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 5103
      },
      {
        "properties": {
          "facing": "west",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 5104
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 5105
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 5106
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 5107
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 5108
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 5109
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 5110
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 5111
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 5112
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 5113
      },
      {
        "properties": {
          "facing": "east",
          "half": "top",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 5114
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "true"
        },
        "id": 5115
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "straight",
          "waterlogged": "false"
        },
        "id": 5116
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "true"
        },
        "id": 5117
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_left",
          "waterlogged": "false"
        },
        "id": 5118
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "true"
        },
        "id": 5119
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "inner_right",
          "waterlogged": "false"
        },
        "id": 5120
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "true"
        },
        "id": 5121
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_left",
          "waterlogged": "false"
        },
        "id": 5122
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "true"
        },
        "id": 5123
      },
      {
        "properties": {
          "facing": "east",
          "half": "bottom",
          "shape": "outer_right",
          "waterlogged": "false"
        },
        "id": 5124
      }
    ]
  },
  "minecraft:command_block": {
    "properties": {
      "conditional": [
        "true",
        "false"
      ],
      "facing": [
        "north",
        "east",
        "south",
        "west",
        "up",
        "down"
      ]
    },
    "states": [
      {
        "properties": {
          "conditional": "true",
          "facing": "north"
        },
        "id": 5125
      },
      {
        "properties": {
          "conditional": "true",
          "facing": "east"
        },
        "id": 5126
      },
      {
        "properties": {
          "conditional": "true",
          "facing": "south"
        },
        "id": 5127
      },
      {
        "properties": {
          "conditional": "true",
          "facing": "west"
        },
        "id": 5128
      },
      {
        "properties": {
          "conditional": "true",
          "facing": "up"
        },
        "id": 5129
      },
      {
        "properties": {
          "conditional": "true",
          "facing": "down"
        },
        "id": 5130
      },
      {
        "properties": {
          "conditional": "false",
          "facing": "north"
        },
        "id": 5131,
        "default": true
      },
      {
        "properties": {
          "conditional": "false",
          "facing": "east"
        },
        "id": 5132
      },
      {
        "properties": {
          "conditional": "false",
          "facing": "south"
        },
        "id": 5133
      },
      {
        "properties": {
          "conditional": "false",
          "facing": "west"
        },
        "id": 5134
      },
      {
        "properties": {
          "conditional": "false",
          "facing": "up"
        },
        "id": 5135
      },
      {
        "properties": {
          "conditional": "false",
          "facing": "down"
        },
        "id": 5136
      }
    ]
  },
  "minecraft:beacon": {
    "states": [
      {
        "id": 5137,
        "default": true
      }
    ]
  },
  "minecraft:cobblestone_wall": {
    "properties": {
      "east": [
        "true",
        "false"
      ],
      "north": [
        "true",
        "false"
      ],
      "south": [
        "true",
        "false"
      ],
      "up": [
        "true",
        "false"
      ],
      "waterlogged": [
        "true",
        "false"
      ],
      "west": [
        "true",
        "false"
      ]
    },
    "states": [
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 5138
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 5139
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "waterlogged": "false",
          "west": "true"
        },
        "id": 5140
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "true",
          "waterlogged": "false",
          "west": "false"
        },
        "id": 5141
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "waterlogged": "true",
          "west": "true"
        },
        "id": 5142
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "waterlogged": "true",
          "west": "false"
        },
        "id": 5143
      },
      {
        "properties": {
          "east": "true",
          "north": "true",
          "south": "true",
          "up": "false",
          "waterlogged": "false",
          "west": "true"
        },
      },
      {