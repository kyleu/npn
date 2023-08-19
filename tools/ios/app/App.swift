// Content managed by Project Forge, see [projectforge.md] for details.
import SwiftUI
import NpnServer

@main
struct Project: App {
    init() {
        print("starting NPN...")
        let path = NSSearchPathForDirectoriesInDomains(.libraryDirectory, .userDomainMask, true)
        let port = NpnServer.CmdLib(path[0])
        print("NPN started on port [\(port)]")
        let url = URL.init(string: "http://localhost:\(port)/")!
        self.cv = ContentView(url: URLRequest(url: url))
    }

    var cv: ContentView

    var body: some Scene {
        WindowGroup {
            cv
        }
    }
}
