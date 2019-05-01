defmodule Collection do
  @moduledoc """
  # Collection
  Contains many functions that interact with list and collections
  """

  @doc """
  Flatten functions was built thinking only with integer list but it works with any element too thanks to elixir language

  ## Examples

      iex> Collection.flatten [1, 2, [3, 4]]
      [1, 2, 3, 4]

      iex> Collection.flatten [1, 2, [3, 4], [5,6, [7, 8, [9]]],0]
      [1, 2, 3, 4, 5, 6, 7, 8, 9, 0]
  """
  def flatten([]), do: []
  def flatten([h|t]), do: flatten(h) ++ flatten(t)
  def flatten(h), do: [h]
end

