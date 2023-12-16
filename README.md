# Killspiel

Killspiel ist eine Anwendung für League of Legends-Twitch-Streamer, um die Zuschauer besser mitfiebern zu lassen.

Es ermöglicht dem Streamer, die Zuschauer über den Ausgang des Spiels "wetten" zu lassen und dies automatisch auswerten
zu lassen und Gewinner festzustellen.
Es kann von verschiedensten Ingame-Werten ausgewählt werden, wie Kills (daher der Name), Tode und KDA aber auch
exotischere Werte wie Gold pro Minute, Pentakills oder BaitPing.

[//]: # (TODO ADD DEMO)
Eine Demo der Oberfläche ist [hier]() (noch nicht) möglich.

### Beispiel, wie es im Chat aussieht

![Screenshot 2023-12-16 174451.png](%7F%2FScreenshot%202023-12-16%20174451.png)

# Installation & erste Nutzung

Unter [Releases](https://github.com/5WlzxEP/Killspiel/releases) die neuste Version für deine Platform herunterladen.

Zum Einrichten ausführen, die Datei enthält alles, was sie zum Laufen benötigt.
Im Falle von Windows wird sich ein Konsolenfenster öffen.
In diesem sind das Log des Programms.
Zur weiteren Einrichtung die Oberfläche aufrufen, welche unter [http://localhost:8088](http://localhost:8088) laufen
sollte.

Von dort aus geht es in den Einstellungen (Zahnrad in der oberen rechten Ecke) weiter.

# Einstellungen

## [settings](http://localhost:8088/settings)

Hier finden sich ein paar allgemeine Einstellungen wie, die Themeauswahl und die Genauigkeit, mit der überprüft wird, ob
eine Schätzung korrekt ist oder nicht.  
Zudem findet sich hier ein Verweis auf die Overlays, welche einfach genutzt werden können.

Von hier aus geht es zu den Guess Collectors & den Daten Collectors

## [Guess Collectors](http://localhost:8088/settings/collector)

Hier lässt sich die Zeit einstellen, die die Zuschauer Zeit haben, um ihre Schätzungen abzugeben.

Zudem wird auf den aktuell einzigen Guess Collector verwiesen, über den Twitchchat.

### [Twitch Chat](http://localhost:8088/settings/collector/twitchchat/)

Hier kannst du alle Einstellungen um die Datensammlung im Twitchchat festlegen.

Benötigt werden:

- ein Twitch Channel, der auf dem gestreamt wird
- ein Api Key, dieser kann über die dritte Seite, über die du mit dem *?* kommst.

Wenn du /announce nutzen möchtest, um Beginn, Ende und Ergebnisse des Killspiels zu verkünden, findest du die
Einstellungen dafür unter dem Advanced Tab.
Den Twitch Access Token kannst du auch als Twitch Api Key nutzen, musst jedoch `oauth:` davor schreiben.

Hinter dem Title, der unter den [Breadcrumbs](https://de.wikipedia.org/wiki/Brotkr%C3%BCmelnavigation) ist, ist ein
Icon, welches angibt, ob die Einstellung korrekt sind.

Weiterhin:

- Prefix: nur Nachrichten, die dieses Prefix haben, werden zu den Schätzungen aufgenommen
- Nachrichten:
  - Begin: wird zum Beginn, genaueres bei Data Collectors
  - Ende: Zeitpunkt, ab dem keine weiteren Schätzungen angenommen werden
  - Auflösen: nach Ermittelung des Ergebnisses und der Sieger, können dies damit verkündet werden

## [Daten Collectors](http://localhost:8088/settings/data)

Hier lassen sich die beiden Daten Collector erreichen.

### [Manuell](http://localhost:8088/settings/data/manual)

Hier lässt sich ein Killspiel manuell starten, beenden und das Ergebnis festlegen.
Bei Killspiel, die manuell gestartet werden, wird das Ergebnis nicht automatisch ermittelt, dies muss ebenfalls manuell
gesetzt werden.
Es kann immer nur ein Killspiel gleichzeitig laufen.

### [Riot Api](http://localhost:8088/settings/data/riot)

Hier finden sich alle Einstellung, die die Riot API betreffen.

Unter Allgemein findet sich die `Häufigkeit der Api-Abfrage`, dies ist die Einstellung in welchem Intervall die Api
befragt wird, ob der Summoner sich im League Spiel befindet.

Unter LoL finden sich die eigentlichen Einstellungen, wie der Riot Api
Key ([Hier erhältlich](https://developer.riotgames.com/)), RiotId, Server und, am wichtigsten, die Kategorie, welche das
Ergebnis darstellt. 
