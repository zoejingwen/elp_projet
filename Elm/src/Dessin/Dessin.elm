module Dessin.Dessin exposing (..)

import Svg exposing (..)
import Svg.Attributes exposing (..)
import Html exposing (..)
import Parser.TcTurtleParser exposing(..)

type alias State =
    { x : Float
    , y : Float
    , angle : Float
    }

degreesToRadians : Float -> Float
degreesToRadians degrees =
    degrees * pi / 180

applyForward : Int -> State -> (State, String)
applyForward distance state =
    let
        radianAngle =
            degreesToRadians state.angle

        newX =
            state.x + toFloat distance * cos radianAngle

        newY =
            state.y - toFloat distance * sin radianAngle

        svgLine =
            String.join ""
                [ "<line x1=\"", String.fromFloat state.x
                , "\" y1=\"", String.fromFloat state.y
                , "\" x2=\"", String.fromFloat newX
                , "\" y2=\"", String.fromFloat newY
                , "\" stroke=\"black\" stroke-width=\"2\" />"
                ]
    in
    ( { state | x = newX, y = newY }, svgLine )

applyLeft : Int -> State -> State
applyLeft angle state =
    { state | angle = state.angle + toFloat angle }

applyRight : Int -> State -> State
applyRight angle state =
    { state | angle = state.angle - toFloat angle }

interpretCommand : Command -> State -> (State, List String)
interpretCommand command state =
    case command of
        Forward distance ->
            let
                (newState, svgLine) =
                    applyForward distance state
            in
            (newState, [ svgLine ])

        Left angle ->
            (applyLeft angle state, [])

        Right angle ->
            (applyRight angle state, [])

        Repeat n commands ->
            let
                repeatHelper (currentState, lines) _ =
                    List.foldl
                        (\cmd (s, acc) ->
                            let
                                (newState, newLines) =
                                    interpretCommand cmd s
                            in
                            (newState, acc ++ newLines)
                        )
                        (currentState, lines)
                        commands
            in
            List.foldl (\_ acc -> repeatHelper acc ()) (state, []) (List.repeat n ())