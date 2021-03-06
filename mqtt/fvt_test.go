/*
 * Copyright (c) 2021 IBM Corp and others.
 *
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v2.0
 * and Eclipse Distribution License v1.0 which accompany this distribution.
 *
 * The Eclipse Public License is available at
 *    https://www.eclipse.org/legal/epl-2.0/
 * and the Eclipse Distribution License is available at
 *   http://www.eclipse.org/org/documents/edl-v10.php.
 *
 * Contributors:
 *    Seth Hoenig
 *    Allan Stockdill-Mander
 *    Mike Robertson
 */

package mqtt

import "os"

// Use setup_IMA.sh for IBM's MessageSight
// Use fvt/rsmb.cfg for IBM's Really Small Message Broker
// Use fvt/mosquitto.cfg for the open source Mosquitto project

var (
	FVTTCP string
	FVTSSL string
)

func init() {
	FVTAddr := os.Getenv("TEST_FVT_ADDR")
	if FVTAddr == "" {
		FVTAddr = "127.0.0.1"
	}
	FVTTCP = "tcp://" + FVTAddr + ":1883"
	FVTSSL = "ssl://" + FVTAddr + ":8883"
}
