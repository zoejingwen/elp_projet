module Parser.TcTurtleParser exposing (..)
import Dessin.Display exposing (..)
type Command
    = Forward Int
    | Left Int
    | Right Int
    | Repeat Int (List Command)
            
parseRight : Parser Command
parseRight =succeed Right
            |. token "Right"
            |. spaces
            |= int 
parseForward : Parser Command
parseForward =succeed Forward
            |. token "Forward"
            |. spaces
            |= int
parseLeft : Parser Command
parseLeft =succeed Left
            |. token "Left"
            |. spaces
            |= int 
parseRepeat : Parser Command
parseRepeat =succeed Repeat
            |. token "Repeat"
            |. spaces
            |= int
            |. spaces
            |= lazy (\_ -> parseCommandList)


parseCommandList : Parser (List Command)
parseCommandList =succeed identify
                |. token "["
                |= parseCommands
                |. token "]"

parseCommands : Parser (List Command)
parseCommands =succeed (\cmd rest -> cmd :: rest)
            |= lazy (\_ -> parseCommand)
            |= lazy (\_ -> parseCommands)
            <|> succeed []

parseCommand : Parser Command
parseCommand =lazy (\_ ->oneOf
            [ parseRight
            , parseForward
            , parseLeft
            , parseRepeat
            ]
    )

