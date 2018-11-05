#!/usr/bin/env python
# -*- coding: utf-8 -*-
import grpc
import metadatos_pb2
import metadatos_pb2_grpc
from cliente import *

# Obtiene datos de Instanciación de un AVIÓN
def IniciarAvion():
    print("Bienvenido al vuelo\n")
    aerolinea = raw_input("[Avion] Ingrese Nombre de la Aerolínea: ")
    codigo = raw_input("[Avion] Ingrese Código de Vuelo: ")
    print("\n* Parámetros de Configuración *\n")
    pesoMax = raw_input("[Avion - "+codigo+"]: Peso Máximo de carga [Kg]: ")
    CapTanque = raw_input("[Avion - "+codigo+"]: Capacidad del tanque de combustible [L]: ")
    ipTorreInicial = raw_input("[Avion - "+codigo+"]: Torre de Control inicial: ")
    datos = {'aerolinea':aerolinea,
            'codigo':codigo,
            'pesoMax':pesoMax,
            'CapTanque':CapTanque,
            'ipTorreInicial':ipTorreInicial}
    return datos

def Despegar(datos,stub):
    raw_input("[Avion - "+datos['codigo']+"]: Presione enter para despegar")
    destino = raw_input("[Avion - "+datos['codigo']+"]: Ingrese destino: ")
    status = comprobar_destino(destino,datos,stub)
    # status = 0 -> TORRE NO TIENE REGISTRADO DESTINO
    # Status = 1 -> DESPEGUE AUTORIZADO
    if (status == 0):
        print("[Avion - "+datos['codigo']+"]: ERROR: Destino desconocido por la Torre\n")
        Despegar(datos,stub)

    elif(status == 1):
        print("[Avion - "+datos['codigo']+"]: Pasando por el Gate...\n")
        print("[Avion - "+datos['codigo']+"]: Todos los pasajeros a bordo y combustible cargado.\n")
        print("[Avion - "+datos['codigo']+"]: Pidiendo instrucciones para despegar...\n")
        statusPermiso = permiso_despegar(destino,datos,stub)
        # statusPermiso = 0 -> NO HAY PISTAS DISPONIBLES
        # statusPermiso = 1 -> PERMISO EXITOSO
        if (statusPermiso == 0):
            print("[Avion - "+datos['codigo']+"]: Todas las pistas están ocupadas, el avión predecesor es "+statusPermiso['predecesor']+"\n")
            statusConfirmacion = esperar_confirmacion(datos,stub) #RESPUESTA DE PISTA Y ALTURA
            print("[Avion - "+datos['codigo']+"]: Pista "+statusConfirmacion['pista']+" asignada y altura de "+statusConfirmacion['altura']+" km.\n")
            print("[Avion - "+datos['codigo']+"]: Despegando...\n")
            datos['ipTorreInicial']=statusConfirmacion['ipdestino']
            return datos

        if (statusPermiso == 1):
            print("[Avion - "+datos['codigo']+"]: Pista "+statusPermiso['pista']+" asignada y altura de "+statusPermiso['altura']+" km.\n")
            print("[Avion - "+datos['codigo']+"]: Despegando...\n")
            datos['ipTorreInicial']=statusPermiso['ipdestino']
            return datos


def Aterrizar(datos):
    stub = Conectar_Torre(datos)
    print("[Avion - "+datos['codigo']+"]: Esperando pista aterrizaje...")
    nPista = contactar_torre(datos,stub)
    print("[Avion - "+datos['codigo']+"]: Aterrizando en la pista "+nPista+"...")
    return stub
    