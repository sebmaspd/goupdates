# Failsafe software updates strategies for firmware

Go code templates to handle failsafe software updates.  

## Real-World Use
In production, the public key is embedded in the firmware or secure storage.  

Firmware updates are signed offline by the us.  

Devices must refuse to install or boot firmware with invalid/missing signatures.  

## Watchdog timer

The Watchdog Timer — a mechanism commonly used in embedded systems to detect and recover from software failures (e.g., a system that hangs or becomes unresponsive).

The watchdog expects a "heartbeat" or "kick" from the running system at regular intervals.

If the system fails to send the heartbeat (due to a crash or hang), the watchdog triggers a reset or recovery.

A watchdog is set with a 3-second timeout.

The simulated system sends a "kick" at irregular intervals.

If the "kick" is late or missing, the watchdog assumes a hang and triggers a reset.

Control the timing logic to simulate hangs, crashes, or recoveries.
