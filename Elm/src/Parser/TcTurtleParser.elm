module Parser.TcTurtleParser exposing (..)

import Parser exposing (..)
import Dessin.Display exposing (..)

type Command
    = Forward Int
    | Left Int
    | Right Int
    | Repeat Int (List Command)

parseRight : Parser Command
parseRight =
    succeed Right
        |. keyword "Right"
        |. spaces
        |= int

parseForward : Parser Command
parseForward =
    succeed Forward
        |. keyword "Forward"
        |. spaces
        |= int

parseLeft : Parser Command
parseLeft =
    succeed Left
        |. keyword "Left"
        |. spaces
        |= int

parseRepeat : Parser Command
parseRepeat =
    succeed Repeat
        |. keyword "Repeat"
        |. spaces
        |= int
        |. spaces
        |= lazy (\_ -> parseCommandList)

parseCommandList : Parser (List Command)
parseCommandList =
    succeed identity
        |. symbol "["
        |= parseCommands
        |. symbol "]"

parseCommands : Parser (List Command)
parseCommands =
    oneOf
        [ succeed (::)
            |= lazy (\_ -> parseCommand)
            |= lazy (\_ -> parseCommands)
        , succeed []
        ]

parseCommand : Parser Command
parseCommand =
    oneOf
        [ parseRight
        , parseForward
        , parseLeft
        , parseRepeat
        ]
