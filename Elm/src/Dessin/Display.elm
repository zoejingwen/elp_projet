module Dessin.Display exposing (..)
import Dessin.Dessin exposing (..)
import Svg exposing (..)
import Html exposing(Html)

display : List Command -> List (Html msg)
display commands =
    let
        initialState =
            { x = 0, y = 0, angle = 0 }

        (finalState, svgLines) =
            List.foldl
                (\cmd (s, acc) ->
                    let
                        (newState, lines) =
                            interpretCommand cmd s
                    in
                    (newState, acc ++ lines)
                )
                (initialState, [])
                commands
    in
    (List.map (\line -> Html.node "g" [] [ Html.text line ]) svgLines)