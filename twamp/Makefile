CXX = g++
CXXFLAGS = -Wall -std=c++11
TARGET = twamp_test
SOURCES = main.cpp twamp.cpp utils.cpp
OBJECTS = $(SOURCES:.cpp=.o)

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CXX) $(CXXFLAGS) -o $@ $^

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $<

clean:
	rm -f $(TARGET) $(OBJECTS)

