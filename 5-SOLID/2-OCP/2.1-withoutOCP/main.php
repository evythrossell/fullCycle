<?php

class Video
{
    private $type;

    public function calcularInteresse() {
        if ($this->type == "Movie") {
            // calcula interesse baseado em filme
        } elseif ($this->type == "TVShow") {
            // calcula interesse baseado em série
        }
    }
}