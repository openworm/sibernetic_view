CC = g++ -std=c++11 -Wall -pedantic -pthread
TARGET:=sview

LIBS := -lGL -lGL -lglut
SRC_DIR := graphics
BUILD_DIR := bin
SRC_EXT := cpp

BINARY_DIR = $(BUILD_DIR)/obj

SRC := $(wildcard $(SRC_DIR)/*.$(SRC_EXT))
OBJ := $(patsubst $(SRC_DIR)/%,$(BINARY_DIR)/%,$(SRC:.$(SRC_EXT)=.o))

CXXFLAGS = $(CC)

all: CXXFLAGS += -O3
all: $(TARGET)

debug: CXXFLAGS += -ggdb -O0
debug: $(TARGET)

$(TARGET): $(OBJ)
	@echo 'Building target: $@'
	$(CXXFLAGS) $(OCL_LIB) -o $(BUILD_DIR)/$(TARGET) $(OBJ) $(LIBS)
	@echo 'Finished building target: $@'

$(BINARY_DIR)/%.o: $(SRC_DIR)/%.$(SRC_EXT)
	@mkdir -p $(BUILD_DIR)
	@mkdir -p $(BINARY_DIR)
	@mkdir -p $(BINARY_DIR)/utils
	@echo 'Building file: $<'
	@echo 'Invoking: GCC C++ Compiler'
	$(CXXFLAGS) $(OCL_INC) -I$(INC_DIR) -c -o "$@" "$<"
	@echo 'Finished building: $<'

clean:
	@echo "Cleaning...";
	$(RM) $(BUILD_DIR)
	$(RM) $(TEST_BIN_DIR)