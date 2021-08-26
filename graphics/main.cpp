//
// Created by serg on 07.04.19.
//

#include <memory>
#include <iostream>

#include "graph.h"
#include "reader.hpp"
#include "model.hpp"

using sibernetic::graphics::graph;

int main(int argc, char** argw){
	auto r = std::make_shared<custom_reader<float>>();
	auto model = new particle_model<float>();
	std::string file_name = "./data/8.txt";
	r->serialize(file_name, model);
	std::cout << model->size() << std::endl;
	std::cout << "x_min " << model->x_min << std::endl;
	std::cout << "x_max " << model->x_max << std::endl;
	std::cout << "y_min " << model->y_min << std::endl;
	std::cout << "y_max " << model->y_max << std::endl;
	std::cout << "z_min " << model->z_min << std::endl;
	std::cout << "z_max " << model->z_max << std::endl;
    graph::model = model;
	graph::run(argc, argw);
}