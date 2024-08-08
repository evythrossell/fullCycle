<?php

interface Movie
{
    public function play();

    public function increaseVolume();
}

class TheLionKing implements Movie
{
    public function play()
    {
        // trecho de código aqui
    }

    public function increaseVolume()
    {
        // trecho de código aqui
    }
}

class ModernTimes implements Movies
{
    public function play()
    {
        // trecho de código aqui
    }

    public function increaseVolume()
    {
       // método não será implementado, pois o filme é mudo 
    }
}