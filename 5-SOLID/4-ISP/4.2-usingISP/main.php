<?php

interface Movie
{
    public function play();
}

interface AudioControl
{
    public function increaseVolume();
}

class TheLionKing implements Movie, AudioControl
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

class ModernTimes implements Movie
{
    public function play()
    {
        // trecho de código aqui
    }
}