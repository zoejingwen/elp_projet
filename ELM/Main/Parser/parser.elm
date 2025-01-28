module TcTurtleParser exposing (..)
import Parser exposing (..)
type program=Right Int
            |Forward Int
            |Left Int
            |Repeat Int
            
