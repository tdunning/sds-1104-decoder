# Siglent Binary Format Decoder

I have an oscilloscope (an [SDS-1104-EX from Siglent](https://siglentna.com/wp-content/uploads/dlm_uploads/2021/08/SDS1000X-ESDS1000X-U_UserManual_EN05B.pdf)) that is very nice to use. One of the things that
convinced me to get it was the WIFI interface and web access. As it turns out, the web interface is 
pretty cranky, but it's still very usable and the great benefit is that you can download 7Msamples 
from four channels with a single click.

The only problem with that is that the data comes in an opaque format which makes it hard to use.
The format is documented (although you do have to make some guesses about some items) by [Siglent](https://siglentna.com/wp-content/uploads/dlm_uploads/2024/06/How-to-Extract-Data-from-the-Binary-File.pdf). Siglent also offers a precompiled binary, but that is a Windows executable that I can't run and that I don't trust anyway.

So I wrote a converter of my own. The nice thing is that in Go, parsing binary formats like this is a breeze.

My converter only handles the analog signals from the SDS-1xxx line of scopes, but it should be very easy to extend to other models if you need that. Raise an issue and we can work together.

# Running the code

To compile and run in one step, try this in the top directory of this program
```shell
go run -o data.csv
```
The input is assumed to come from a file named `usr_wf_data.bin`, but you can add a different name to the command above if you renamed the data from default filename that the scope gives you.

The output of this looks like this:
```
time,channel3,channel4
0,0.24000000298023225,0.8000000035762787
0,0.26000000298023224,0.8000000035762787
2e-08,0.26000000298023224,0.7200000035762787
4e-08,0.26000000298023224,0.7200000035762787
6e-08,0.28000000298023225,0.7600000035762787
8e-08,0.24000000298023225,0.8000000035762787
```
There is one column for the tiem and one column for each channel you are using. The values are resolved to seconds (for time) and volts (for the measurements)

Here's an example of some signals I was analyzing:

<img width="891" height="655" alt="image" src="https://github.com/user-attachments/assets/8fc5f042-ff40-4180-af1e-baa90b323752" />

The Julia code that produced this from the CSV files is just this
```julia
using CSV, DataFrames, Plots, FFTW, Statistics

"""
Reads a CSV file and computes the spectrum.
"""
function readData(fname)
  x = CSV.read(fname, DataFrame)
  x[:,:diff] = x[!,:channel3] - x[!,:channel4]
  df = 1/mean(diff(x[!,:time])) / size(x,1)
  w = sin.(LinRange(0,π,size(x,1))) .^ 2
  sp = fft(w .* x[:,:diff])
  return x, sp, df
end

x1, sp1, df1 = readData("data1.csv")
x1[:,:diff] = x1[:,:channel3] - x1[:,:channel4]
x2, sp2, df2 = readData("data2.csv")
x1[:,:diff] = x1[:,:channel3] - x1[:,:channel4]

plot(x1[1:10000, :time]*1e6, x1[1:10000, :diff])
plot!(x2[1:10000, :time]*1e6, x2[1:10000, :diff])
plot!(xlab="Time (μs)", ylab="Volts")
```


