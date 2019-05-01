# Elixir Flatten

<a name="installation"></a>

## Installation

First you need to install [Elixir](https://elixir-lang.org/install.html).

Second is highly recommended to add Elixir’s bin path to your PATH environment variable to ease development.

On Windows, there are [instructions for different versions](http://www.computerhope.com/issues/ch000549.htm) explaining the process.

On Unix systems, you need to [find your shell profile file](https://unix.stackexchange.com/a/117470/101951), and then add to the end of this file the following line reflecting the path to your Elixir installation:

```
export PATH="$PATH:/path/to/elixir/bin"
```

Then you can clone.

```
$ git clone git@github.com:iver/flatten.git
```

## Execution

Go to elixir folder...

```
$ cd flatten/elixir/
$ iex 
Erlang/OTP 21 [erts-10.2.3] [source] [64-bit] [smp:4:4] [ds:4:4:10] [async-threads:1] [hipe] [dtrace]

Interactive Elixir (1.8.1) - press Ctrl+C to exit (type h() ENTER for help)
```

Now you can load the file...

```
iex(1)> c "flatten.ex"
[Collection]
```
And invoke it...

```
iex(2)> Collection.flatten [1, 2, [3, 4], [5,6, [7, 8, [9]]],0]
[1, 2, 3, 4, 5, 6, 7, 8, 9, 0]
```

## Test

If you wish to run the test:

```
$ elixir flatten_test.exs
```