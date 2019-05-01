if !System.get_env("EXERCISM_TEST_EXAMPLES") do
  Code.load_file("flatten.ex", __DIR__)
end

ExUnit.start()
ExUnit.configure(exclude: :pending, trace: true)

defmodule WordsTest do
  use ExUnit.Case

  test "one list" do
    list = [1, 2, [3, 4]]
    result = Collection.flatten list
    expected = [1, 2, 3, 4]
    assert result == expected
  end

  test "many list" do
    list = [[1], 2, [[3,4], 5], [[[]]], [[[6]]], 7, 8, []]
    result = Collection.flatten list 
    expected = [1, 2, 3, 4, 5, 6, 7, 8]
    assert result == expected 
  end
  test "repeated elements" do
    list = [[1], 2, 5, 6, [[3,4], 5], [[[]]], [[[6]]], 7, 8, []]
    result = Collection.flatten list 
    expected = [1, 2, 5, 6, 3, 4, 5, 6, 7, 8]
    assert result == expected 
  end
end
