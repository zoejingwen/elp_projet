module Main exposing (..)

import Browser
import Html exposing (Html, Attribute, button, div, input, text)
import Html.Attributes exposing (..)
import Html.events exposing (onClick, onInput)
import Parser exposing (..)

-- MAIN
main =
    Browser.sandbox { init = init, update = update, view = view }

-- MODEL


-- UPDATE
