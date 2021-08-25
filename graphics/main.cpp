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
	std::string file_name = "./data/1.txt";
	r->serialize(file_name, model);
	std::cout << model->size() << std::endl;
	graph::run(argc, argw);
}