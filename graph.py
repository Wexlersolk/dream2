import json
import matplotlib.pyplot as plt
from datetime import datetime

# Define the file path
file_path = '/home/wexlersolk/work/dream2/file/dream.json'

# Function to read JSON file and extract dates and scores
def read_json(file_path):
    with open(file_path, 'r') as file:
        data = json.load(file)

    dates = []
    scores = []

    for entry in data:
        # Convert date string to datetime object for proper plotting
        date = datetime.fromisoformat(entry['Date'].replace("Z", "+00:00"))
        dates.append(date)
        scores.append(entry['Score'])

    return dates, scores

# Function to create the line graph
def create_line_graph(dates, scores):
    plt.figure(figsize=(10, 5))
    plt.plot(dates, scores, marker='o')
    plt.title('Score over Time')
    plt.xlabel('Date')
    plt.ylabel('Score')
    plt.xticks(rotation=45)
    plt.grid()
    plt.tight_layout()  # Adjust layout to prevent clipping of tick-labels
    plt.show()  # Show the graph in a new window

# Function to create a bar graph with scores vs dates
def create_bar_graph(dates, scores):
    plt.figure(figsize=(10, 5))
    plt.bar(dates, scores, width=0.4, color='skyblue', edgecolor='black')
    plt.title('Score per Date')
    plt.xlabel('Date')
    plt.ylabel('Score')
    plt.xticks(rotation=45)
    plt.grid(axis='y')
    plt.tight_layout()
    plt.show()  # Show the bar graph in a new window

# Main function to execute the program
def main():
    dates, scores = read_json(file_path)
    create_line_graph(dates, scores)  # Show the line graph
    create_bar_graph(dates, scores)    # Show the bar graph

if __name__ == '__main__':
    main()

