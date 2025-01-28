module Main exposing (..)

import Browser
import Html exposing (Html, Attribute, button, div, input, text)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick, onInput)
-- import Parser.TcTurtleParser exposing (..)

-- MAIN

main =
    Browser.sandbox { init = init, update = update, view = view }

-- MODEL

type alias Model =
    { commandes : String
    }

init : Model
init =
    Model ""

-- UPDATE

type Msg
    = Start 

update : Msg -> Model -> Model
update msg model =
    case msg of
        Start -> model -- display model.commandes -- adapter display



-- VIEW

view : Model -> Html Msg
view model =
    div []
        [ div [] [text "Type in your TcTurtle code below :"]
        , input [ placeholder "exemple : [Repeat 360 [ Right 1, Forward 1]]", value model.commandes ] []
        , button [ onClick Start] [ text "Draw"]
        ] -- ajouter le display svg