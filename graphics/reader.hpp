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
                if(cur_line.empty()) {
                    continue;
                }
                cur_line.erase(std::remove(cur_line.begin(), cur_line.end(), '\r'),cur_line.end()); // crlf win fix
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
                                model->y_max = static_cast<T>(stod(cur_line));
                                break;
                            case 3:
                                model->y_min = static_cast<T>(stod(cur_line));
                                break;
                            case 4:
                                model->z_max = static_cast<T>(stod(cur_line));
                                break;
                            case 5:
                                model->z_min = static_cast<T>(stod(cur_line));
                                break;
                        }
                        ++param_line_cnt;
                        continue;
                    }
                    case POS: {
                        particle<T> p;
                        auto data = *get_vector(cur_line);
                        auto pos = std::array<T, 4>{data[0],data[1],data[2],data[3]};
                        auto vel = std::array<T, 4>{data[4],data[5],data[6],data[7]};
                        p.velocity = vel;
                        p.position = pos;
                        p.type = static_cast<int>(p.position[3]);
                        p.particle_id = index;
                        model->push_back(p);
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
    std::shared_ptr <std::array<T, 8>> get_vector(const std::string &line) {
        std::shared_ptr <std::array<T, 8>> v(new std::array<T, 8>());
        std::stringstream ss(line);
        ss >> (*v)[0] >> (*v)[1] >> (*v)[2] >> (*v)[3] >> (*v)[4] >> (*v)[5] >> (*v)[6] >> (*v)[7]; // TODO check here!!!
        return v;
    }
};
#endif //SIBERNETIC_CUSTOM_READER_HPP
