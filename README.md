| ch1_on | 0x00-0x03 | on/off status of CH1, 1 - ON, 0 - OFF 32-bit signed integer.|
| ch2_on | 0x04-0x07 | on/off status of CH2, 1 - ON, 0 - OFF 32-bit integer |
| ch3_on | 0x08-0x0b | on/off status of CH3, 1 - ON, 0 - OFF 32-bit integer |
| ch4_on | 0x0c-0x0f | on/off status of CH4, 1 - ON, 0 - OFF 32-bit integer |
| ch1_volt_div_val | 0x10-0x1f | V/div value of CH1, such as 2.48 mV/div. Unit of value, such as V from 0x1c-0x1f, refer to Table for the details. Units of value’s magnitude (MICRO) from 0x18-0x1b, refer to Table for the details. 64-bit float point, data of value from 0x10-0x17. |
| ch2_volt_div_val | 0x20-0x2f | V/div value of CH2, such as 2.48 mV/div. Unit of value, such as V from 0x2c-0x2f, refer to Table for the details. Units of value’s magnitude (MICRO) from 0x28-0x2b, refer to Table for the details. 64-bit float point, data of value from 0x20-0x27. |
| ch3_volt_div_val | 0x30-0x3f | V/div value of CH3, such as 2.48 mV/div. Unit of value, such as V from 0x3c-0x3f, refer to Table for the details. Units of value’s magnitude (MICRO) from 0x38-0x3b, refer to Table for the details. 64-bit float point, data of value from 0x30-0x37. |
| ch4_volt_div_val | 0x40-0x4f | V/div value of CH4, such as 2.48 mV/div. Unit of value, such as V from 0x4c-0x4f, refer to Table for the details. Units of value’s magnitude (MICRO) from 0x48-0x4b, refer to Table for the details. 64-bit float point, data of value from 0x40-0x47. |
| ch1_vert_offset | 0x50-0x5f | Offset value of CH1, such as 2.48 mV. Unit of value, such as V from 0x5c-0x5f, refer to Table for the details. Units of value’s magnitude (MICRO) from 0x58-0x5b, refer to Table for the details. 64-bit float point, data of value from 0x50-0x57. |
| ch2_vert_offset | 0x60-0x6f | Offset value of CH2, such as 2.48 mV. Unit of value, such as V from 0x6c-0x6f, refer to Table for the details. Units of value’s magnitude (MICRO) from 0x68-0x6b, refer to Table for the details. 64-bit float point, data of value from 0x60-0x67. |
| ch3_vert_offset | 0x70-0x7f | Offset value of CH3, such as 2.48 mV. Unit of value, such as V from 0x7c-0x7f, refer to Table for the details. Units of value’s magnitude (MICRO) from 0x78-0x7b, refer to Table for the details. 64-bit float point, data of value from 0x70-0x77. |
| ch4_vert_offset | 0x80-0x8f | Offset value of CH4, such as 2.48 mV. Unit of value, such as V from 0x8c-0x8f, refer to Table for the details. Units of value’s magnitude (MICRO) from 0x88-0x8b, refer to Table for the details. 64-bit float point, data of value from 0x80-0x87. |
| digital_on | 0x90-0x93 | on/off status of digital, 1 - ON, 0 - OFF 32-bit integer |
| d0_d15_on | 0x94-0xd3 | on/off status of d0-d15, 1 - ON, 0 - OFF 32-bit integer |
| time_div | 0xd4-0xe3 | Time div (time base) value, Such as 2.48 ms/div. Unit of value, such as s from 0xe0-0xe3, refer to Table for the details. Units of value’s magnitude (MICRO) from 0xdc-0xdf, refer to Table for the details. 64-bit float point, data of value from 0xd4-0xdb. |
| time_delay | 0xe4-0xf3 | Time delay (Trigger delay) value, Such as 2.48 ms. Unit of value, such as s from 0xf0-0xf3, refer to Table for the details. Units of value’s magnitude (MICRO) from 0xec-0xef, refer to Table for the details. 64-bit float point, data of value from 0xe4-0xeb
| wave_length | 0xf4-0xf7 | Wave length of the data points for analog channel. 32-bit integer |
| Sample_rate | 0xf8-0x107 | Sample Rate value for analog channel, Such as 500M Sa/s. Unit of value, such as Sa from 0x104-0x107, refer to Table for the details. Units of value’s magnitude (MEGA) from 0x100-0x103, Refer to Table for the details. 64-bit float point, data of value from 0xf8-0xff. |
| digital_wave_length | 0x108-0x10b | Wave length of the data points for 32-bit integer |
| digital_sample_rate | 0x10c-0x11b | Sample Rate value for digital, Such as 500M Sa/s. Unit of value, such as Sa from 0x118-0x11b, refer to Table for the details. Units of value’s magnitude (MEGA) from 0x114-0x117, Refer to Table for the details. 64-bit float point, data of value from 0x10c-0x113.
| reserved | 0x11c~ | reserved |
| reserved | ~0x7ff | reserved |
| Wave_data | 0x800-end | Data from CH1 to D15. Only data of the enabled channel(s) are stored to the file.
