import os
import pandas as pd


dir = "data/unformatted/"

# list and open each file
for filename in os.listdir(dir):
    with open('data/unformatted/'+ filename, 'r') as f:
        coordinates = []
        reading_coordinates = False

        matrix = []

        lines = f.readlines()
        for line in lines:
            line = line.strip()
            if line == "NODE_COORD_SECTION":
                reading_coordinates = True
                continue
            if line == "EOF":
                break
            if reading_coordinates:
                parts = line.split()
                coordinates.append((float(parts[1]), float(parts[2])))

        for c0 in coordinates:
            distances_lines = []
            for c1 in coordinates:
               d = ((c0[0] - c1[0])**2 + (c0[1] - c1[1])**2)**0.5
               distances_lines.append(d)
            matrix.append(distances_lines)

        df = pd.DataFrame(matrix)

        fn = filename.split('.')
        df.to_csv("data/formatted/" + fn[0] + ".csv", header=False, index=False)
