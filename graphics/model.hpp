#ifndef MODEL_HPP
#define MODEL_HPP

#include <vector>

#include "particle.hpp"

template<class T = float> struct particle_model {
    T x_min;
    T x_max;
    T y_min;
    T y_max;
    T z_min;
    T z_max;
    int particle_id;
    std::vector<particle<T>> container;
    void push_back(particle<T> &p) {
        container.push_back(p);
    }
    particle<T>& get_particle(int index) {
        return container.at(index);
    }
};

#endif