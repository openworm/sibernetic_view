#ifndef PARTICLE_HPP
#define PARTICLE_HPP

#include <array>

#define ARRAY_DIM 4

template<class T=float>struct particle
{
    std::array<T, ARRAY_DIM> position;
    std::array<T, ARRAY_DIM> velocity;
    int particle_id;
};


#endif