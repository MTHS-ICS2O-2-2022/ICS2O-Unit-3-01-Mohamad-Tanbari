// Copyright (c) 2020 Mohamad All rights reserved
//
// Created by: Mohamad
// Created on: Sep 2020
// This file contains the JS functions for index.html

/*
* This function 
*/
function calculate () {
  // debugging statements
  console.log('button clicked')

  // grab input from the text box
  const aBase = parseInt(document.getElementById('a-base').value)
  const bBase = parseInt(document.getElementById('b-base').value)
  const height = parseInt(document.getElementById('height').value)

  // calculate the area
  const area = (aBase + bBase) / 2 * height

  // output the area
  document.getElementById('answer').innerHTML = "The area of the Trapezoid is: " + area + " cm²"
}
