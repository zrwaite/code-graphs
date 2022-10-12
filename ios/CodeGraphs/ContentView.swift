//
//  ContentView.swift
//  CodeGraphs
//
//  Created by Zachary Waite on 2022-10-10.
//

import SwiftUI

func AuthorizeWakatime() {
    print("ello")
}


struct ContentView: View {
    var body: some View {
        ZStack {
            Color.black.ignoresSafeArea()
            VStack{
                Text("Waka Widgets")
                    .font(.largeTitle)
                    .modifier(WWText())
                Text("Welcome!")
                    .font(.title)
                    .modifier(WWText())
                Button("Authorize Wakatime", action: AuthorizeWakatime)
                Link("Learn SwiftUI", destination: URL(string: "https://www.hackingwithswift.com/quick-start/swiftui")!)
            }
            .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .top)
        }
    }
}
