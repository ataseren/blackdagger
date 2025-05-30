.. _Configuration Options:

Configuration Options
=====================

.. contents::
    :local:

.. _Environment Variables:

Environment Variables
----------------------

The following environment variables can be used to configure the BLACKDAGGER. Default values are provided in the parentheses:

- ``BLACKDAGGER_HOST`` (``127.0.0.1``): The host to bind the server to.
- ``BLACKDAGGER_PORT`` (``8080``): The port to bind the server to.
- ``BLACKDAGGER_DAGS`` (``$HOME/.config/blackdagger/dags``): The directory containing the DAGs.
- ``BLACKDAGGER_IS_BASICAUTH`` (``0``): Set to 1 to enable basic authentication.
- ``BLACKDAGGER_BASICAUTH_USERNAME`` (``""``): The username to use for basic authentication.
- ``BLACKDAGGER_BASICAUTH_PASSWORD`` (``""``): The password to use for basic authentication.
- ``BLACKDAGGER_LOG_DIR`` (``$HOME/.local/share/blackdagger/logs``): The directory where logs will be stored.
- ``BLACKDAGGER_DATA_DIR`` (``$HOME/.local/share/blackdagger/history``): The directory where application data will be stored.
- ``BLACKDAGGER_SUSPEND_FLAGS_DIR`` (``$HOME/.config/blackdagger/suspend``): The directory containing DAG suspend flags.
- ``BLACKDAGGER_ADMIN_LOG_DIR`` (``$HOME/.local/share/admin``): The directory where admin logs will be stored.
- ``BLACKDAGGER_BASE_CONFIG`` (``$HOME/.config/blackdagger/base.yaml``): The path to the base configuration file.
- ``BLACKDAGGER_NAVBAR_COLOR`` (``""``): The color to use for the navigation bar. E.g., ``red`` or ``#ff0000``.
- ``BLACKDAGGER_NAVBAR_TITLE`` (``BLACKDAGGER``): The title to display in the navigation bar. E.g., ``BLACKDAGGER - PROD`` or ``BLACKDAGGER - DEV``
- ``BLACKDAGGER_WORK_DIR``: The working directory for DAGs. If not set, the default value is DAG location. Also you can set the working directory for each DAG steps in the DAG configuration file. For more information, see :ref:`specifying working dir`.
- ``BLACKDAGGER_WORK_DIR``: The working directory for DAGs. If not set, the default value is DAG location. Also you can set the working directory for each DAG steps in the DAG configuration file. For more information, see :ref:`specifying working dir`.
- ``BLACKDAGGER_CERT_FILE``: The path to the SSL certificate file.
- ``BLACKDAGGER_KEY_FILE`` : The path to the SSL key file.
- ``BLACKDAGGER_SKIP_INITIAL_DAG_PULLS``: Set to 1 to skip the automatic pull of default DAG YAML files during startup.
- ``BLACKDAGGER_DAG_REPO``: The prefix for the DAG repository. This is used to pull the categorized DAG YAML files from the repository. The default value is ``https://github.com/ErdemOzgen/blackdagger-``.

Note: If ``BLACKDAGGER_HOME`` environment variable is not set, the default value is ``$HOME/.config/blackdagger``.

Config File
--------------

You can create ``config.yaml`` file in the ``$BLACKDAGGER_HOME`` directory (default: ``$HOME/.config/blackdagger``) to override the default configuration values. The following configuration options are available:

.. code-block:: yaml

    host: <hostname for web UI address>                          # default: 127.0.0.1
    port: <port number for web UI address>                       # default: 8080

    # path to the DAGs directory
    dags: <the location of DAG configuration files>              # default: ${BLACKDAGGER_HOME}/dags
    
    # Web UI Color & Title
    navbarColor: <ui header color>                               # header color for web UI (e.g. "#ff0000")
    navbarTitle: <ui title text>                                 # header title for web UI (e.g. "PROD")
    
    # Basic Auth
    isBasicAuth: <true|false>                                    # enables basic auth
    basicAuthUsername: <username for basic auth of web UI>       # basic auth user
    basicAuthPassword: <password for basic auth of web UI>       # basic auth password

    # API Token
    isAuthToken: <true|false>                                    # enables API token
    authToken: <token for API access>                            # API token

    # Base Config
    baseConfig: <base DAG config path>                           # default: ${BLACKDAGGER_HOME}/config.yaml

    # Working Directory
    workDir: <working directory for DAGs>                        # default: DAG location

    # DAG Configuration
    skipInitialDAGPulls: <true|false>                            # skip initial DAG pulls
    dagRepo: <DAG repository prefix>                             # default: https://github.com/ErdemOzgen/blackdagger-

    # SSL Configuration
    tls:
        certFile: <path to SSL certificate file>
        keyFile: <path to SSL key file>

.. _Host and Port Configuration:

Server's Host and Port Configuration
-------------------------------------

To specify the host and port for running the BLACKDAGGER server, there are a couple of ways to do it.

The first way is to specify the ``BLACKDAGGER_HOST`` and ``BLACKDAGGER_PORT`` environment variables. For example, you could run the following command:

.. code-block:: sh

    BLACKDAGGER_PORT=8000 blackdagger server

The second way is to use the ``--host`` and ``--port`` options when running the ``blackdagger server`` command. For example:

.. code-block:: sh

    blackdagger server --port=8000

See :ref:`Environment Variables` for more information.

.. _base configuration:

Base Configuration for all DAGs
---------------------------------

Creating a base configuration (default path: ``~/.config/blackdagger/base.yaml``) is a convenient way to organize shared settings among all DAGs. The path to the base configuration file can be configured. See :ref:`Configuration Options` for more details.

Example:

.. code-block:: yaml

    # directory path to save logs from standard output
    logDir: /path/to/stdout-logs/

    # history retention days (default: 30)
    histRetentionDays: 3

    # Email notification settings
    mailOn:
      failure: true
      success: true

    # SMTP server settings
    smtp:
      host: "smtp.foo.bar"
      port: "587"
      username: "<username>"
      password: "<password>"

    # Error mail configuration
    errorMail:
      from: "foo@bar.com"
      to: "foo@bar.com"
      prefix: "[Error]"

    # Info mail configuration
    infoMail:
      from: "foo@bar.com"
      to: "foo@bar.com"
      prefix: "[Info]"