# Killspiel

Killspiel ist eine Anwendung für League of Legends-Twitch-Streamer, um die Zuschauer besser mitfiebern zu lassen.

Es ermöglicht dem Streamer, die Zuschauer über den Ausgang des Spiels "wetten" zu lassen und dies automatisch auswerten
zu lassen und Gewinner festzustellen.
Es kann von verschiedensten Ingame-Werten ausgewählt werden, wie Kills (daher der Name), Tode und KDA aber auch
exotischere Werte wie Gold pro Minute, Pentakills oder BaitPing.

Eine Demo der Oberfläche ist [hier](https://5wlzxep.github.io/Killspiel/) vorhanden.

### Beispiel, wie es im Chat aussieht

![Screenshot 2023-12-16 174451.png](%7F%2FScreenshot%202023-12-16%20174451.png)

# Installation & erste Nutzung

Unter [Releases](https://github.com/5WlzxEP/Killspiel/releases) die neuste Version für deine Platform herunterladen.

Zum Einrichten ausführen, die Datei enthält alles, was sie zum Laufen benötigt.
Im Falle von Windows wird sich ein Konsolenfenster öffen.
In diesem sind das Log des Programms.
Zur weiteren Einrichtung die Oberfläche aufrufen, welche unter [http://localhost:8088](http://localhost:8088) laufen
sollte.

# Voraussetzung

Es wird ein Riot-Api-Key benötigt. Dieser kann [hier](https://developer.riotgames.com/) angefragt werden. Zum Ausprobieren reicht ein Development-Key aus, dieser muss aber alle 24h erneuert werden. Bei längerer Nutzung ist ein [Personal Api Key](https://developer.riotgames.com/app-type) empfohlen.