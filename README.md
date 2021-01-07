# Lab 3 - Sistemas Distribuidos

Equipo : FalopApps
Nombre : Benjamín Molina
Rol : 201573005-9

Nombre : Tomás Escalona
Rol : 201573031-8

## Instrucciones previas a la ejecución

Ingresar en terminal:

cd broker-service
./generate.sh
make

## Ejecución - Iniciar un Servidor DNS

Ingresar en terminal desde la carpeta broker-service:

./DNS --ip <DIRECCIÓN_IP_DE_ESTE_SERVIDOR> --number <NÚMERO_DEL_SERVIDOR_DNS>

## Ejecución - Iniciar un Broker

Ingresar en terminal desde la carpeta broker-service:

./Broker --broker*ip <DIRECCIÓN_IP_DEL_BROKER> --dns1_ip <DIRECCIÓN_IP_DEL_SERVIDOR_DNS*#1> --dns2*ip <DIRECCIÓN_IP_DEL_SERVIDOR_DNS*#2> --dns3*ip <DIRECCIÓN_IP_DEL_SERVIDOR_DNS*#3> --dns_port <PUERTO_DE_LOS_SERVIDORES_DNS>

## Ejecución - Iniciar un Admin

Ingresar en terminal desde la carpeta broker-service:

./Admin --dns_port <PUERTO_DE_LOS_SERVIDORES_DNS> --broker_addr <DIRECCIÓN_IP_CON_PUERTO_QUE_USARA_EL_BROKER>

## Ejecución - Iniciar un Cliente

./Client --broker_addr <DIRECCIÓN_IP_CON_PUERTO_QUE_USARA_EL_BROKER>

Comandos adicionales:

Los ejecutables cuentan con más argumentos adicionales, para mayor información usar --help, por ej:

./DNS --help
