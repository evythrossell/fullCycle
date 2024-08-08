<?php

abstract class Video 
{
    abstract public function calcularInteresse()
}

class Movie extends Video
{
    public function calcularInteresse()
    {
        // trecho de código aqui
    }
}

class TVShow extends Video {
    public function calcularInteresse()
}