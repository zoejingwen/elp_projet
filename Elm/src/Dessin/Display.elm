module Dessin.Display exposing (..)

import Dessin.Dessin exposing (..)
import Svg exposing (..)
import Svg.Attributes exposing (..)

display : List Command -> Svg msg
display commands =
    let
        initialState =
            { x = 250, y = 250, angle = 0 } -- Position initiale au centre de l'écran

        (_, svgLines) =
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
    svg
        [ width "500"
        , height "500"
        , viewBox "0 0 500 500"
        ]
        (List.map (\line -> g [] [ text_ [] [ text line ] ]) svgLines)
