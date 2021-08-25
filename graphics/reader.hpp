//
// Created by sergey on 12.02.19.
//

#ifndef SIBERNETIC_CUSTOM_READER_HPP
#define SIBERNETIC_CUSTOM_READER_HPP

#include "error.h"
#include <fstream>
#include <iostream>
#include <regex>

#include "particle.hpp"
#include "model.hpp"

template<class T>
class custom_reader{
    enum LOADMODE {
        NOMODE = -1, PARAMS, MODEL, POS, VEL
    };
public:
    void serialize(const std::string &file_name, particle_model<T> *model) {
        std::ifstream file(file_name.c_str(), std::ios_base::binary);
        LOADMODE mode = NOMODE;
        int index = 0;
        int param_line_cnt = 0;
        if (file.is_open()) {
            while (file.good()) {
                std::string cur_line;
                std::getline(file, cur_line);
                cur_line.erase(std::remove(cur_line.begin(), cur_line.end(), '\r'),
                               cur_line.end()); // crlf win fix
                auto i_space = cur_line.find_first_not_of(" ");
                auto i_tab = cur_line.find_first_not_of("\t");
                if (i_space) {
                    cur_line.erase(cur_line.begin(), cur_line.begin() + i_space);
                }
                if (i_tab) {
                    cur_line.erase(cur_line.begin(), cur_line.begin() + i_tab);
                }
                if (cur_line == "bbox") {
                    mode = PARAMS;
                    continue;
                }
                if (cur_line == "position") {
                    mode = POS;
                    continue;
                } else if (cur_line == "velocity") {
                    mode = VEL;
                    continue;
                }
                switch (mode) {
                    case PARAMS: {
                    // Super Dumb
                        switch(param_line_cnt) {
                            case 0:
                                model->x_max = static_cast<T>(stod(cur_line));
                                break;
                            case 1:
                                model->x_min = static_cast<T>(stod(cur_line));
                                break;
                            case 2:
                                model->y_min = static_cast<T>(stod(cur_line));
                                break;
                            case 3:
                                model->y_max = static_cast<T>(stod(cur_line));
                                break;
                            case 4:
                                model->z_min = static_cast<T>(stod(cur_line));
                                break;
                            case 5:
                                model->z_max = static_cast<T>(stod(cur_line));
                                break;
                        }
                        ++param_line_cnt;
                        continue;
                    }
                    case POS: {
                        particle<T> p;
                        p.position = *get_vector(cur_line);
                        p.type = static_cast<int>(p.position[3]);
                        model->push_back(p);
                        break;
                    }
                    case VEL: {
                        model->get_particle(index).velocity = *get_vector(cur_line);
                        model->get_particle(index).particle_id = index;
                        ++index;
                        break;
                    }
                    default: {
                        break;
                    }
                }
            }
        } else {
            throw "Check your file name or path there is no file with name " + file_name;
        }
        file.close();
    }

private:
    std::shared_ptr <std::array<T, 4>> get_vector(const std::string &line) {
        std::shared_ptr <std::array<T, 4>> v(new std::array<T, 4>());
        std::stringstream ss(line);
        ss >> (*v)[0] >> (*v)[1] >> (*v)[2] >> (*v)[3]; // TODO check here!!!
        return v;
    }
};
#endif //SIBERNETIC_CUSTOM_READER_HPP
