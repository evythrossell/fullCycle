<?php

class Course 
{
    private $name;
    private $category;
    private $description;

    public function connection()
    {
        $pdo = new PDO();
        return $pdo;
    }

    public function createCategory()
    {
        $this->connection()->insert($this->category);
    }

    public function createCourse() 
    {
        $this->connection()->insert($this->name);
    }

    public function getName()
    {
        return $this->name;
    }

    public function setName($name)
    {
        $this->name = $name
        return $this;
    }

    public function getCategory()
    {
        return $this->category;
    }

    public function setCategory($category)
    {
        return $this->category = $category
        return $this;
    }
}