<p align="center">
  <img src="./assets/images/blackdaggerReadme.png" width="500" alt="blackdagger-logo">
</p>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/ErdemOzgen/blackdagger">
    <img src="https://goreportcard.com/badge/github.com/ErdemOzgen/blackdagger" />
  </a>
  <a href="https://codecov.io/gh/ErdemOzgen/blackdagger">
    <img src="https://codecov.io/gh/ErdemOzgen/blackdagger/branch/main/graph/badge.svg?token=CODZQP61J2" />
  </a>
  <a href="https://github.com/erdemozgen/blackdagger/releases">
    <img src="https://img.shields.io/github/release/erdemozgen/blackdagger.svg" />
  </a>
  <a href="https://godoc.org/github.com/ErdemOzgen/blackdagger">
    <img src="https://godoc.org/github.com/ErdemOzgen/blackdagger?status.svg" />
  </a>
  <img src="https://github.com/ErdemOzgen/blackdagger/actions/workflows/test.yaml/badge.svg" />
</p>

<div align="center">

[Installation](#installation) | [Quick Start Docs](https://blackdagger.readthedocs.io/en/latest/quickstart.html#launch-the-web-ui) 

</div>

<h1><b>Blackdagger</b></h1>

Blackdagger represents a significant advancement in automation technology, offering a comprehensive solution for orchestrating complex workflows in DevOps, DevSecOps, MLOps, MLSecOps, and Continuous Red Teaming (CART) environments.

At its core, Blackdagger simplifies the management and execution of intricate workflows through its user-friendly approach and powerful functionality. Leveraging a declarative YAML format, Blackdagger enables users to define automation pipelines using a Directed Acyclic Graph (DAG), facilitating clear and concise expression of task dependencies and execution logic.

What sets Blackdagger apart is its simplicity and versatility. Unlike traditional cron-based schedulers or workflow orchestration platforms, Blackdagger eliminates the need for extensive scripting or coding. With a built-in Web UI, users can easily manage, rerun, and monitor automation pipelines in real-time, streamlining the workflow management process. Additionally, Blackdagger offers native Docker support, enabling seamless integration with containerized environments, and a versatile toolset for task execution, including making HTTP requests and executing commands over SSH.

**What Sets Blackdagger Apart?**

1. **Declarative YAML Format**: Blackdagger simplifies workflow definition with a declarative YAML format, allowing users to specify command dependencies using a Directed Acyclic Graph (DAG). This intuitive approach makes it easy to define complex workflows and manage task dependencies without the need for extensive scripting or coding.
2. **Web UI for Visual Management**: With its built-in Web UI, Blackdagger provides users with a visually intuitive interface for managing, rerunning, and monitoring automation pipelines. Users can easily track the real-time status of workflows, view execution logs, and make configuration changes directly from their browser, eliminating the need for manual intervention.
3. **Native Docker Support**: Blackdagger natively supports Docker container management, enabling seamless integration with containerized environments. Users can run arbitrary Docker containers as part of their automation workflows, making it easy to orchestrate tasks across distributed infrastructure and microservices architectures.
4. **Versatile Task Execution**: From making HTTP requests to executing commands over SSH, Blackdagger offers a versatile toolset for task execution. Whether it's interacting with external APIs, running custom code snippets, or managing infrastructure components, Blackdagger empowers users to automate a wide range of tasks with ease.

## **Highlights**
- Single binary file installation
- Declarative YAML format for defining DAGs
- Web UI for visually managing, rerunning, and monitoring pipelines
- Use existing programs without any modification
- Self-contained, with no need for a DBMS
- Suitable for Continuous Red Teaming (CART)
- Suitable for DevOps and DevSecOps
- Suitable for MLOps and MLSecOps

You can find everything about Blackdagger, including this README, in our [documentation](https://blackdagger.readthedocs.io).

## Installation

**Via Bash script**

```
curl -L https://raw.githubusercontent.com/ErdemOzgen/blackdagger/main/scripts/downloader.sh | bash
```

**Via Docker**

```
# in blackdagger repository
docker compose up
```

**Via GitHub Release Page**

Download the latest binary from the [Releases page](https://github.com/ErdemOzgen/blackdagger/releases) and place it in your `$PATH` (e.g. `/usr/local/bin`).


**Quick Start Guide**

1. Launch the Web UI

Start the server and scheduler with the command `blackdagger start-all` and browse to `http://127.0.0.1:8080` to explore the Web UI.

2. Create a New DAG

Navigate to the DAG List page by clicking the menu in the left panel of the Web UI. Then create a DAG by clicking the `New DAG` button at the top of the page. Enter `example` in the dialog.

*Note: DAG (YAML) files will be placed in `~/.blackdagger/dags` by default. See [Configuration Options](https://blackdagger.readthedocs.io/en/latest/config.html) for more details.*

3. Edit the DAG

Go to the `SPEC` Tab and hit the `Edit` button. Copy & Paste the following example and click the `Save` button.

Example:

```
schedule: "* * * * *" # Run the DAG every minute
steps:
  - name: s1
    command: echo Hello blackdagger
  - name: s2
    command: echo done!
    depends:
      - s1
```

4. Execute the DAG

You can execute the example by pressing the `Start` button. You can see "Hello blackdagger" in the log page in the Web UI.


## **Features**

- Web User Interface
- Command Line Interface (CLI) with several commands for running and managing DAGs
- YAML format for defining DAGs, with support for various features including:
  - Execution of custom code snippets
  - Parameters
  - Command substitution
  - Conditional logic
  - Redirection of stdout and stderr
  - Lifecycle hooks
  - Repeating task
  - Automatic retry
- Executors for running different types of tasks:
  - Running arbitrary Docker containers
  - Making HTTP requests
  - Sending emails
  - Running jq command
  - Executing remote commands via SSH
- Email notification
- Scheduling with Cron expressions
- REST API Interface
- Basic Authentication over HTTPS


## **Use Cases**

- **Data Pipeline Automation:** Schedule ETL tasks for data processing and centralization.
- **Infrastructure Monitoring:** Periodically check infrastructure components with HTTP requests or SSH commands.
- **Automated Reporting:** Generate and send periodic reports via email.
- **Batch Processing:** Schedule batch jobs for tasks like data cleansing or model training.
- **Task Dependency Management:** Manage complex workflows with interdependent tasks.
- **Microservices Orchestration:** Define and manage dependencies between microservices.
- **CI/CD Integration:** Automate code deployment, testing, and environment updates.
- **Alerting System:** Create notifications based on specific triggers or conditions.
- **Custom Task Automation:** Define and schedule custom tasks using code snippets.
- **Model Training Automation:** Automate the training of machine learning models by scheduling jobs that run on new data sets. Use Blackdagger to manage dependencies between data preprocessing, training, evaluation, and deployment tasks.
- **Model Deployment Pipeline:** Create a DAG to automate the deployment of trained models to production environments, including steps for model validation, containerization with Docker, and deployment using SSH commands.
- **Security Scans Integration:** Schedule regular security scans and static code analysis as part of the CI/CD pipeline. Use Blackdagger to orchestrate these tasks, ensuring that deployments are halted if vulnerabilities are detected.
- **Automated Compliance Checks:** Set up workflows to automatically run compliance checks against infrastructure and codebase, reporting results via HTTP requests to compliance monitoring tools.
- **Automated Penetration Testing:** Schedule and manage continuous penetration testing activities. Define dependencies in Blackdagger to ensure that penetration tests are conducted after deployment but before wide release, using Docker containers to isolate testing environments.
- **Threat Simulation and Response:** Automate the execution of threat simulations to test the effectiveness of security measures. Use Blackdagger to orchestrate complex scenarios involving multiple steps, such as breaching a system, escalating privileges, and exfiltrating data, followed by automated rollback and alerting.

## Usage

### **Web UI**

The easiest and most efficient way to use Blackdagger is using its Web UI. Below, there are some features and functions of Web UI that you can benefit.

#### DAG Details

It shows the real-time status, logs, and DAG configurations. You can edit DAG configurations on a browser.

  ![example](assets/images/demo.png)

  You can switch to the vertical graph with the button on the top right corner.

  ![Details-TD](assets/images/ui-details.png?raw=true)

#### DAGs List

It shows all DAGs and the real-time status.

  ![DAGs](assets/images/ui-dags.png)

#### Search DAGs

It greps given text across all DAGs.
  ![History](assets/images/ui-search.png?raw=true)

#### Execution History

It shows past execution results and logs.

  ![History](assets/images/ui-history.png)

#### DAG Execution Log

It shows the detail log and standard output of each execution and step.

  ![DAG Log](assets/images/ui-logoutput.png)


### **CLI Commands**

Blackdagger also supports CLI commands to perform all possible actions that you can take on Web UI. Below, you can see all CLI commands and their usage examples.

```bash
# Runs the DAG
blackdagger start [--params=<params>] <file>

# Displays the current status of the DAG
blackdagger status <file>

# Re-runs the specified DAG run
blackdagger retry --req=<request-id> <file>

# Stops the DAG execution
blackdagger stop <file>

# Restarts the current running DAG
blackdagger restart <file>

# Dry-runs the DAG
blackdagger dry [--params=<params>] <file>

# Launches both the web UI server and scheduler process
blackdagger start-all [--host=<host>] [--port=<port>] [--dags=<path to directory>]

# Launches the blackdagger web UI server
blackdagger server [--host=<host>] [--port=<port>] [--dags=<path to directory>]

# Starts the scheduler process
blackdagger scheduler [--dags=<path to directory>]

# Shows the current binary version
blackdagger version
```

## **Example Workflow**

This example workflow showcases a data pipeline typically implemented in DevOps and Data Engineering scenarios. It demonstrates an end-to-end data processing cycle starting from data acquisition and cleansing to transformation, loading, analysis, reporting, and ultimately, cleanup.

![Details-TD](assets/images/example.png)

The YAML code below represents this workflow:

```yaml
# Environment variables used throughout the pipeline
env:
  - DATA_DIR: /data
  - SCRIPT_DIR: /scripts
  - LOG_DIR: /log
  # ... other variables can be added here

# Handlers to manage errors and cleanup after execution
handlerOn:
  failure:
    command: "echo error"
  exit:
    command: "echo clean up"

# The schedule for the workflow execution in cron format
# This schedule runs the workflow daily at 12:00 AM
schedule: "0 0 * * *"

steps:
  # Step 1: Pull the latest data from a data source
  - name: pull_data
    command: "bash"
    script: |
      echo `date '+%Y-%m-%d'`
    output: DATE

 # Step 2: Cleanse and prepare the data
  - name: cleanse_data
    command: echo cleansing ${DATA_DIR}/${DATE}.csv
    depends:
      - pull_data

  # Step 3: Transform the data
  - name: transform_data
    command: echo transforming ${DATA_DIR}/${DATE}_clean.csv
    depends:
      - cleanse_data

  # Parallel Step 1: Load the data into a database
  - name: load_data
    command: echo loading ${DATA_DIR}/${DATE}_transformed.csv
    depends:
      - transform_data

  # Parallel Step 2: Generate a statistical report
  - name: generate_report
    command: echo generating report ${DATA_DIR}/${DATE}_transformed.csv
    depends:
      - transform_data

  # Step 4: Run some analytics
  - name: run_analytics
    command: echo running analytics ${DATA_DIR}/${DATE}_transformed.csv
    depends:
      - load_data

  # Step 5: Send an email report
  - name: send_report
    command: echo sending email ${DATA_DIR}/${DATE}_analytics.csv
    depends:
      - run_analytics
      - generate_report

  # Step 6: Cleanup temporary files
  - name: cleanup
    command: echo removing ${DATE}*.csv
    depends:
      - send_report
```

## **Contribution**

blackdagger is a single command line tool that uses the local file system to store data, so no database management system or cloud service is required. DAGs are defined in a declarative YAML format, and existing programs can be used without modification.

----

Feel free to contribute in any way you want! Share ideas, questions, submit issues, and create pull requests. Check out our [Contribution Guide](https://blackdagger.readthedocs.io/en/latest/contrib.html) for help getting started.

We welcome any and all contributions!

## **License**

This project is licensed under the GNU GPLv3. It was forked from [Dagu](https://github.com/dagu-dev/dagu) and has been adapted to serve a different purpose. While Dagu is an excellent project, its current objectives do not align with ours.